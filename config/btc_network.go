package config

import (
	"github.com/btcsuite/btcd/chaincfg"
)

var BtcNetwork *chaincfg.Params

func InitBtcNetwork(ginMode string) {
	//if ginMode == gin.ReleaseMode {
	BtcNetwork = &chaincfg.MainNetParams
	//} else {
	//	BtcNetwork = &chaincfg.TestNet3Params
	//}
}
