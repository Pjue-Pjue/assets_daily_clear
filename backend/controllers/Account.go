package controllers

import (
	"Pjue-Pjue/assets_daily_clear_web/backend/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pborman/uuid"
	"time"
)

func AddAccount(c *gin.Context) {
	data := struct {
		Market     string `json:"market"`
		AccessKey  string `json:"access_key"`
		SecretKey  string `json:"secret_key"`
		PassPhrase string `json:"pass_phrase"`
		Label      string `json:"label"`
	}{}
	if c.BindJSON(&data) != nil {
		c.JSON(200, gin.H{
			"code": 403,
			"msg":  "params error",
		})
		return
	}

	// 检查是否存在
	exists, err := models.AccountExists(data.Market, data.AccessKey, data.SecretKey, data.PassPhrase)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 403,
			"msg":  "internal error",
		})
		return
	}

	if exists {
		c.JSON(200, gin.H{
			"code": 403,
			"msg":  "already exist",
		})
		return
	}

	unixTime := time.Now().Unix()
	account := models.Account{
		Id:         uuid.New(),
		Market:     data.Market,
		AccessKey:  data.AccessKey,
		SecretKey:  data.SecretKey,
		UpdateAt:   unixTime,
		Status:     1,
		PassPhrase: data.PassPhrase,
		Label:      data.Label,
	}
	err = models.AddAccount(account)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 403,
			"msg":  "write internal error",
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
	})
}

func GetAccounts(c *gin.Context) {
	list, err := models.GetAccounts()
	if err != nil {
		c.JSON(200, gin.H{
			"code": 403,
			"msg":  "internal error",
		})
		fmt.Println(err.Error())
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": list,
	})
}

func UpdateAccount(c *gin.Context) {
	data := struct {
		Ids    []string `json:"ids"`
		Status int      `json:"status"`
	}{}
	if c.BindJSON(&data) != nil {
		c.JSON(200, gin.H{
			"code": 403,
			"msg":  "params error",
		})
		return
	}
	if len(data.Ids) <= 0 {
		c.JSON(200, gin.H{
			"code": 403,
			"msg":  "params error",
		})
		return
	}
	for _, v := range data.Ids {
		s, err := models.GetAccountByID(v)
		if err != nil {
			c.JSON(200, gin.H{
				"code": 403,
				"msg":  "internal error",
			})
			fmt.Println(err.Error())
			return
		}

		if s.Status == data.Status {
			continue
		}

		err = models.UpdateById(v, data.Status)
		if err != nil {
			c.JSON(200, gin.H{
				"code": 403,
				"msg":  err.Error(),
			})
			return
		}
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "操作成功",
	})
}
