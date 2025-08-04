package orm

import (
	"cmc-server/models"
	"encoding/json"
	"os"
)

func InitRbacData() error {
	err := InitPromission()

	if err != nil {
		return err
	}

	return nil
}

func InitPromission() error {

	_, err := Engine.Exec("TRUNCATE TABLE promission")

	if err != nil {
		return err
	}

	data, err := os.ReadFile("static/promission.json")
	if err != nil {
		return err
	}

	var promissionList []models.Promission

	err = json.Unmarshal(data, &promissionList)

	if err != nil {
		return err
	}

	for _, v := range promissionList {
		_, err = Engine.Insert(&v)
		if err != nil {
			return err
		}
	}

	return nil
}
