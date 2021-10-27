package utils

import (
	"channelServe/models"
	"errors"
)

func SendTo(from, to, amount int, channelId string) error {
	playerFrom, err := models.GetPlayerById(from,channelId)
	if err != nil {
		return err
	}
	playerTo, err := models.GetPlayerById(to,channelId)
	if err != nil {
		return err
	}
	if playerFrom.Credit < amount {
		return errors.New("Insufficient Balance")
	}
	playerFrom.Credit -= amount
	err = models.UpdatePLayer(playerFrom, channelId)
	if err != nil {
		return err
	}
	playerTo.Credit += amount
	err = models.UpdatePLayer(playerTo, channelId)
	if err != nil {
		return err
	}
	return nil
}
func ExitChannel(channelId string, uid int) error {
	err := models.DeletePlayer(uid, channelId)
	if err != nil {
		return err

	}
	return nil
}
