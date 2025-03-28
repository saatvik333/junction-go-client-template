package junction

// "github.com/airchains-network/junction/x/zksequencer/types"

// func RegisterStation(account cosmosaccount.Account, addr string, client cosmosclient.Client, msg types.MsgInitStation) {
// 	ctx := context.Background()

// 	txRes, err := client.BroadcastTx(ctx, account, &msg)
// 	if err != nil {
// 		utils.Log.Error("Register collection failed")
// 		fmt.Println(err)
// 		return
// 	}
// 	utils.Log.Info(txRes.TxHash)
// 	// events := txRes.Events
// 	// var collectionId string
// 	// for _, event := range events {
// 	// 	if event.Type == "collection_registered" {
// 	// 		utils.Log.Success("Collection registered successfully")
// 	// 		for _, attr := range event.Attributes {

// 	// 			// logs.Log.Debug(attr.Key)
// 	// 			// logs.Log.Success(attr.Value)
// 	// 			if attr.Key == "collection_id" {
// 	// 				collectionId = attr.Value
// 	// 				logs.Log.Success(fmt.Sprintf("Collection ID: %s", collectionId))
// 	// 			}
// 	// 		}
// 	// 	}
// 	// }
// 	// logs.Log.Success("http://0.0.0.0:1317/airchains-network/junction/vrf/fetch_collection/" + collectionId)
// 	// logs.Log.Success("http://0.0.0.0:1317/airchains-network/junction/vrf/fetch_collection_details/" + collectionId)
// 	// return collectionId, txRes.Code == 0
// }
