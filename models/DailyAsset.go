package models

type DailyAsset struct {
	Id        string  `db:"id, primarykey"`
	AccountId string  `db:"account_id"`
	UpdateAt  int64   `db:"update_at"`
	AssetUsd  float64 `db:"asset_usd"`
}

func AddDailyAsset(item DailyAsset) error {
	err := DB.Insert(&item)
	return err
}

func GetDailyRecordByAccountId(accountId string) ([]DailyAsset, error) {
	var ret []DailyAsset
	_, err := DB.Select(&ret, "select * from daily_asset_usd_valuate where account_id=? order by update_at asc", accountId)
	if err != nil {
		return ret, err
	}
	return ret, nil
}