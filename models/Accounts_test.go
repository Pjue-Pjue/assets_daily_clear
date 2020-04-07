package models

import (
	"github.com/pborman/uuid"
	"testing"
	"time"
)

func TestAddAccount(t *testing.T) {
	item := Account{
		Id:        uuid.New(),
		Market:    "OKEX",
		AccessKey: "2234235",
		SecretKey: "fdhdjdesj",
		UpdateAt:  time.Now().Unix(),
		Status:    1,
	}
	err := AddAccount(item)
	if err != nil {
		t.Log(err)
	}
}

func TestGetDailyRecordByAccountId(t *testing.T) {
	//item := DailyAsset{
	//	Id: uuid.New(),
	//	AccountId: "1",
	//	AssetUsd: 2.004,
	//	UpdateAt: time.Now().Unix(),
	//}
	//err := AddDailyAsset(item)
	//if err != nil {
	//	t.Log(err.Error())
	//	return
	//}
	res, err := GetDailyRecordByAccountId("1")
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(len(res))
}