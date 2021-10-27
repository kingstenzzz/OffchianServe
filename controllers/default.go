package controllers

import (
	"channelServe/models"
	"channelServe/utils"
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"strconv"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.html"

}

func (c *MainController) Join() error {
	user := c.GetString("addr")
	channel := c.GetString("channel")
	deposit, err := c.GetInt("deposit")
	fmt.Printf("user:%v,channel:%v", user, channel)
	player := models.Player{Addr: user, Credit: deposit, Withdrawn: 0, Deposit: deposit}
	uid, err := models.AddPlayer(player, channel)
	var result map[string]string
	if err != nil {
		result = map[string]string{"code": "404", "message": "failed", "err": err.Error()}
	} else {
		result = map[string]string{"uid": strconv.Itoa(uid), "code": "200", "message": "成功"}
	}
	c.Data["json"] = result
	//定义返回json

	c.ServeJSON()
	return nil
}
func (c *MainController) Send() error {
	toId, err := c.GetInt("to")
	amount, err := c.GetInt("amount")
	uid, err := c.GetInt("uid")
	channelId := c.GetString("channel")
	fmt.Printf("send:%v,to:%v", amount, toId)
	//player, err := models.GetPlayerById(uid)
	err = utils.SendTo(uid, toId, amount, channelId)
	var result map[string]string
	if err != nil {
		result = map[string]string{"code": "404", "message": "failed", "err": err.Error()}
	} else {
		result = map[string]string{"uid": strconv.Itoa(uid), "code": "200", "message": "send successfully"}
	}
	c.Data["json"] = result
	//定义返回json
	c.ServeJSON()
	return nil
}

func (c *MainController) Exit() error {
	uid, err := c.GetInt("uid")
	if err != nil {
		return err
	}
	channelId := c.GetString("channel")
	err = utils.ExitChannel(channelId, uid)

	var result map[string]string
	if err != nil {
		result = map[string]string{"code": "404", "message": "failed", "err": err.Error()}
	} else {
		result = map[string]string{"uid": strconv.Itoa(uid), "code": "200", "message": "delete successfully"}
	}
	c.Data["json"] = result
	//定义返回json
	c.ServeJSON()
	return nil

}

func (c *MainController) Show() {
	fmt.Println("showInfo")
	addr := c.GetString("playerAddr")
	channelId := c.GetString("channel")
	player, err := models.GetPlayerByAddr(addr, channelId)

	if err != nil {
		result := map[string]string{"code": "404", "message": "failed", "err": err.Error()}
		c.Data["json"] = result
		c.ServeJSON()
	} else {
		playerByte, err := json.Marshal(player)
		if err !=nil{
			result := map[string]string{"code": "404", "message": "failed", "err": err.Error()}
			c.Data["json"] = result
			c.ServeJSON()
		}else {
			var  response map[string]interface{}
			err = json.Unmarshal(playerByte, &response)
			c.Data["json"] = response
			c.ServeJSON()
		}

	}

}

func (c *MainController) CreateChanel() {
	channelID := c.GetString("channelId")
	err := models.CreateChannel(channelID)
	if err != nil {
	}
	var result map[string]string
	if err != nil {
		result = map[string]string{"code": "404", "message": "failed", "err": err.Error()}
	} else {
		result = map[string]string{"channel": channelID, "code": "200", "message": "成功"}
	}
	c.Data["json"] = result
	//定义返回json
	c.ServeJSON()
}

func (c *MainController) PlayerPost() {
	function := c.GetString("action")
	switch function {
	case "join":
		c.Join()
	case "send":
		c.Send()
	case "exit":
		c.Exit()
	case "dispute":
	}

}

func (c *MainController) pLayerGet() {
	function := c.GetString("action")
	switch function {
	case "join":
	case "send":
	case "exit":
	case "dispute":
	}

}
