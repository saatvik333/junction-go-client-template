package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"

	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/group/edwards25519"
)

func NewKeyPair() (privateKeyX kyber.Scalar, publicKeyX kyber.Point) {
	// Initialize the Kyber suite for Edwards25519 curve
	suite := edwards25519.NewBlakeSHA256Ed25519()

	// Generate a new private key
	privateKey := suite.Scalar().Pick(suite.RandomStream())

	// Derive the public key from the private key
	publicKey := suite.Point().Mul(privateKey, nil)

	return privateKey, publicKey
}

func GenerateVRFProof(suite kyber.Group, privateKey kyber.Scalar, data []byte, nonce int64) ([]byte, []byte, error) {
	// Convert nonce to a deterministic scalar
	nonceBytes := big.NewInt(nonce).Bytes()
	nonceScalar := suite.Scalar().SetBytes(nonceBytes)

	// Generate proof like in a Schnorr signature: R = g^k, s = k + e*x
	R := suite.Point().Mul(nonceScalar, nil) // R = g^k
	hash := sha256.New()
	rBytes, _ := R.MarshalBinary()
	hash.Write(rBytes)
	hash.Write(data)
	e := suite.Scalar().SetBytes(hash.Sum(nil))                             // e = H(R||data)
	s := suite.Scalar().Add(nonceScalar, suite.Scalar().Mul(e, privateKey)) // s = k + e*x

	// The VRF output (pseudo-random value) is hash of R combined with data
	vrfHash := sha256.New()
	vrfHash.Write(rBytes)         // Incorporate R
	vrfHash.Write(data)           // Incorporate input data
	vrfOutput := vrfHash.Sum(nil) // This is the deterministic "random" output

	// Serialize R and s into the proof
	sBytes, _ := s.MarshalBinary()
	proof := append(rBytes, sBytes...)

	return proof, vrfOutput, nil
}

func LoadHexPrivateKey(hexPrivateKey string) (privateKey kyber.Scalar, err error) {
	privateKeyBytes, err := hex.DecodeString(hexPrivateKey)
	if err != nil {
		fmt.Printf("Error decoding private key: %v\n", err)
		return nil, err
	}
	suite := edwards25519.NewBlakeSHA256Ed25519()
	privateKey = suite.Scalar().SetBytes(privateKeyBytes)
	return privateKey, nil
}

type RequestCommitmentV2Plus struct {
	BlockNum         uint64
	StationId        string
	UpperBound       uint64
	RequesterAddress string
}

func SerializeRequestCommitmentV2Plus(rc RequestCommitmentV2Plus) ([]byte, error) {
	var buf bytes.Buffer

	// Encode the blockNum
	err := binary.Write(&buf, binary.BigEndian, rc.BlockNum)
	if err != nil {
		return nil, fmt.Errorf("failed to encode blockNum: %w", err)
	}

	// Encode the stationId as a fixed size or prefixed with its length
	// Here, we choose to prefix with length for simplicity
	if err := binary.Write(&buf, binary.BigEndian, uint64(len(rc.StationId))); err != nil {
		return nil, fmt.Errorf("failed to encode stationId length: %w", err)
	}
	buf.WriteString(rc.StationId)

	// Encode the upperBound
	err = binary.Write(&buf, binary.BigEndian, rc.UpperBound)
	if err != nil {
		return nil, fmt.Errorf("failed to encode upperBound: %w", err)
	}

	// Encode the requesterAddress as a fixed size or prefixed with its length
	if err := binary.Write(&buf, binary.BigEndian, uint64(len(rc.RequesterAddress))); err != nil {
		return nil, fmt.Errorf("failed to encode requesterAddress length: %w", err)
	}
	buf.WriteString(rc.RequesterAddress)

	// Encode the extraArgs
	//buf.WriteByte(rc.ExtraArgs)

	return buf.Bytes(), nil
}

func VerifyVRFProof(hexPublicKey string, serializedRC []byte, proof []byte, vrfOutput []byte) (bool, error) {
	suite := edwards25519.NewBlakeSHA256Ed25519()

	publicKey, err := LoadHexPublicKey(hexPublicKey)
	if err != nil {
		return false, fmt.Errorf("error loading public key: %w", err)
	}

	// Deserialize R and s from the proof
	pointSize := suite.Point().MarshalSize()
	R, s := suite.Point(), suite.Scalar()
	if err := R.UnmarshalBinary(proof[:pointSize]); err != nil {
		return false, fmt.Errorf("error unmarshalling R: %w", err)
	}
	s.SetBytes(proof[pointSize:])

	// Recompute e = H(R||data) from the proof and data
	hash := sha256.New()
	rBytes, _ := R.MarshalBinary()
	hash.Write(rBytes)
	hash.Write(serializedRC)
	e := suite.Scalar().SetBytes(hash.Sum(nil))

	// Verify the equation R == g^s * y^-e
	gs := suite.Point().Mul(s, nil) // g^s, correct usage

	// Correct calculation for y^e where 'y' is publicKey (a point) and 'e' is a scalar.
	ye := suite.Point().Mul(e, publicKey)

	yeInv := suite.Point().Neg(ye)            // -y^e, correct usage
	expectedR := suite.Point().Add(gs, yeInv) // g^s * y^-e, correct combination

	if !R.Equal(expectedR) {
		return false, fmt.Errorf("invalid VRF proof")
	}

	// Verify the VRF output matches the hash of R and data
	vrfHash := sha256.New()
	vrfHash.Write(rBytes)
	vrfHash.Write(serializedRC)
	expectedVrfOutput := vrfHash.Sum(nil)
	if !bytes.Equal(vrfOutput, expectedVrfOutput) {
		return false, fmt.Errorf("invalid VRF output")
	}

	return true, nil
}

// LoadHexPublicKey loads a public key from a hexadecimal string
func LoadHexPublicKey(hexPublicKey string) (kyber.Point, error) {
	// Decode the hexadecimal string to a byte slice
	publicKeyBytes, err := hex.DecodeString(hexPublicKey)
	if err != nil {
		return nil, fmt.Errorf("error decoding public key: %w", err)
	}

	// Initialize the Kyber suite for the Edwards25519 curve
	suite := edwards25519.NewBlakeSHA256Ed25519()

	// Convert the byte slice into a Kyber point representing the public key
	publicKey := suite.Point()
	if err := publicKey.UnmarshalBinary(publicKeyBytes); err != nil {
		return nil, fmt.Errorf("error unmarshalling public key: %w", err)
	}

	return publicKey, nil
}
