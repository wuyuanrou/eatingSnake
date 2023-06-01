package dao

import (
	"eating_snake/model"
	"eating_snake/tool"
	"fmt"
	"log"
	"strconv"
)

//跨包调用函数名必须大写
func GetBestHistoryDataByMode(mode int) (historys []model.HistoryData, total int, err error) {
	var sqlStr string
	i := 1

	err = DB.QueryRow("SELECT count(*) FROM history WHERE mode=?", mode).Scan(&total)
	fmt.Printf("共%d条数据\n", total)

	if total <= 30 {
		sqlStr = fmt.Sprintf("SELECT length FROM history WHERE mode = %d ORDER BY length desc", mode) //只取长度
	} else {
		sqlStr = fmt.Sprintf("SELECT length FROM history WHERE mode = %d LIMIT %d OFFSET %d ORDER BY length desc", mode, 30, total-30)
	}

	rows, err := DB.Query(sqlStr)
	for rows.Next() {
		history := model.HistoryData{}
		history.Num = i
		history.Mode = tool.ModeToModeName(strconv.Itoa(mode))
		err := rows.Scan(&history.Length)
		if err != nil {
			log.Fatal("dao getBestHistoryDataByMode err:", err)
		}
		historys = append(historys, history)
		i += 1
	}
	_ = rows.Close()
	return historys, total, nil
}

func GetBestHistoryData() (historys []model.HistoryData, total int, err error) {
	var sqlStr string
	i := 1

	err = DB.QueryRow("SELECT count(*) FROM history").Scan(&total)
	fmt.Printf("共%d条数据\n", total)

	if total <= 30 {
		sqlStr = fmt.Sprintf("SELECT length, mode FROM history ORDER BY length desc") //只取长度
	} else {
		sqlStr = fmt.Sprintf("SELECT length, mode FROM history  LIMIT %d OFFSET %d ORDER BY length desc", 30, total-30)
	}

	rows, err := DB.Query(sqlStr)
	for rows.Next() {
		history := model.HistoryData{}
		history.Num = i
		err := rows.Scan(&history.Length, &history.Mode)
		if err != nil {
			log.Fatal("dao getBestHistoryData err:", err)
		}
		historys = append(historys, history)
		i += 1
	}
	_ = rows.Close()
	return historys, total, nil
}

func GetRecentHistoryData() (historys []model.HistoryData, total int, err error) {
	var sqlStr string

	err = DB.QueryRow("SELECT count(*) FROM history").Scan(&total)
	fmt.Printf("共%d条数据\n", total)

	if total <= 30 {
		sqlStr = fmt.Sprintf("SELECT length, mode FROM history") //只取长度
	} else {
		sqlStr = fmt.Sprintf("SELECT length, mode FROM history LIMIT %d OFFSET %d", 30, total-30)
	}

	rows, err := DB.Query(sqlStr)
	for rows.Next() {
		history := model.HistoryData{}
		history.Num = 0
		err := rows.Scan(&history.Length, &history.Mode)
		if err != nil {
			log.Fatal("dao getHistoryData err:", err)
		}
		historys = append(historys, history)
	}
	_ = rows.Close()
	return historys, total, nil
}

func InsertHistoryData(length int, mode int) error {
	_, err := DB.Exec("INSERT INTO history (length, mode) VALUES (?, ?)", length, mode)
	if err != nil {
		fmt.Printf("dao插入成绩错误, error:[%v]", err.Error())
		return err
	}
	//num, _ := result.RowsAffected()
	//fmt.Printf("dao插入成绩成功, affected rows:[%d]\n", num)
	return nil
}
