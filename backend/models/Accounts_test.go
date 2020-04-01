package models

import (
	"github.com/pborman/uuid"
	"testing"
	"time"
)

func TestAddAccount(t *testing.T) {
	item := Account{
		Id: uuid.New(),
		Market: "OKEX",
		AccessKey: "2234235",
		SecretKey: "fdhdjdesj",
		UpdateAt: time.Now().Unix(),
		Status: 1,
	}
	err := AddAccount(item)
	if err != nil {
		t.Log(err)
	}
}
