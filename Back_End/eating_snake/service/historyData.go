package service

import (
	"eating_snake/dao"
	"eating_snake/model"
	"eating_snake/tool"
	"fmt"
	"log"
)

func GetBestHistoryDataByMode(modeName string) (historys []model.HistoryData, err error) {
	mode := tool.ModeNameToMode(modeName)

	historys, _, err = dao.GetBestHistoryDataByMode(mode)
	if err != nil {
		log.Fatal("service getBestHistoryDataByMode err:", err.Error())
	}

	return historys, nil
}

func GetBestHistoryData() (historys []model.HistoryData, err error) {
	raw_historys, _, err := dao.GetBestHistoryData()
	if err != nil {
		log.Fatal("service getBestHistoryData err:", err.Error())
	}

	for _, history := range raw_historys {
		history.Mode = tool.ModeToModeName(history.Mode)
		historys = append(historys, history)
	}
	return historys, nil
}

func GetRecentHistoryData() (historys []model.HistoryData, err error) {
	raw_historys, total, err := dao.GetRecentHistoryData()
	if err != nil {
		log.Fatal("service getHistoryData err:", err.Error())
	}

	j := 1
	for i := total - 1; i >= 0; i-- {
		raw_historys[i].Num = j
		raw_historys[i].Mode = tool.ModeToModeName(raw_historys[i].Mode)
		historys = append(historys, raw_historys[i])
		j += 1
	}

	return historys, nil
}

func InsertHistoryData(length int, modeName string) error {
	mode := tool.ModeNameToMode(modeName)
	err := dao.InsertHistoryData(length, mode)
	if err != nil {
		fmt.Println("service插入成绩错误, error:", err.Error())
		return err
	}
	return nil
}
