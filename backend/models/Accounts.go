package models

import "fmt"

type Account struct {
	Id         string `db:"id, primarykey"`
	Market     string `db:"market"`
	AccessKey  string `db:"access_key"`
	SecretKey  string `db:"secret_key"`
	UpdateAt   int64  `db:"update_at"`
	Status     int    `db:"status"`
	PassPhrase string `db:"pass_phrase"`
	Label      string `db:"label"`
}

func AccountExists(market, accessKey, secretKey, passPhrase string) (bool, error) {
	n, err := DB.SelectInt("select count(*) from accounts where market=? and access_key=? and secret_key=? and pass_phrase=? and status!=0",
		market, accessKey, secretKey, passPhrase)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return n > 0, nil
}

func AddAccount(item Account) error {
	err := DB.Insert(&item)
	return err
}

func GetAccounts() ([]Account, error) {
	var ret []Account
	_, err := DB.Select(&ret, "select * from accounts where status!=0 order by update_at desc")
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func GetAccountByID(id string) (Account, error) {
	var ret Account
	err := DB.SelectOne(&ret, "select * from accounts where id=? and status!=0", id)
	if err != nil {
		return ret, err
	}
	return ret, nil
}

func UpdateById(id string, status int) error {
	_, err := DB.Exec("update accounts set status=? where id=?", status, id)
	return err
}