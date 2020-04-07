package main

import (
	"Pjue-Pjue/assets_daily_clear_web/models"
	"Pjue-Pjue/dc-exchanges-api/okex"
	"errors"
	"fmt"
	"github.com/jasonlvhit/gocron"
	"github.com/pborman/uuid"
	"time"
)

func init() {

}

func main() {
	gocron.Every(1).Day().At("17:30").Do(dailyClear)
	go gocron.Start()
	select {}
}

func dailyClear() {
	taskItems, err := models.GetRunningAccounts()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, v := range taskItems {
		account, err := GetAccountInfo(v.AccessKey, v.SecretKey, v.PassPhrase, v.Market)
		if err != nil {
			continue
		}
		err = DailyRecord(account, v.Id)
		if err != nil {
			fmt.Printf("DailyRecord err: %v", err.Error())
			break
		}
	}
}

func GetAccountInfo(apikey, secretKey, passphrase string, market string) (u float64, err error) {
	switch market {
	case "OKEX":
		u = GerOkexTotalAsset(apikey, secretKey, passphrase)
		break
	default:
		err = errors.New(fmt.Sprintf("暂不支持:%v", market))
		break
	}
	return
}

func GerOkexTotalAsset(apikey, secretKey, passphrase string) (u float64) {
	c := okex.NewAccountOKEXClient(apikey, secretKey, passphrase)
	var err error
	var asset okex.AssetValuation
	for {
		asset, err = c.GetAssetValuation(0, "USD")
		if err != nil {
			time.Sleep(10 * time.Second)
			continue
		}
		break
	}
	u = asset.Balance
	return
}

func DailyRecord(u float64, id string) error {
	item := models.DailyAsset{
		Id:        uuid.New(),
		AccountId: id,
		AssetUsd:  u,
		UpdateAt:  time.Now().Unix(),
	}
	err := models.AddDailyAsset(item)
	return err
}
