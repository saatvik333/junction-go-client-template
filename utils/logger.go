package utils

import (
	"go.uber.org/zap"
)

var Log *zap.Logger

func InitLogger() {
	var err error
	Log, err = zap.NewProduction() // You can use zap.NewDevelopment() for a development-friendly logger
	if err != nil {
		panic(err)
	}
}
