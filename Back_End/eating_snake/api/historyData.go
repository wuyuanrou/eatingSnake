package api

import (
	"eating_snake/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetBestHistoryDataByMode(c *gin.Context) {
	modeName := c.Param("modeName")
	fmt.Println("当前模式:", modeName)
	if modeName == "" {
		msg := fmt.Sprintf("api-GetHistoryDataByMode-获取mode字段失败")
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  msg,
		})
		return
	}

	historys, err := service.GetBestHistoryDataByMode(modeName)
	if err != nil {
		msg := fmt.Sprintf("api-GetHistoryData-获取历史战绩错误")
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  msg,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": historys,
	})
}

func GetBestHistoryData(c *gin.Context) {
	historys, err := service.GetBestHistoryData()
	if err != nil {
		msg := fmt.Sprintf("api-GetBestHistoryData-获取最佳战绩错误")
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  msg,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": historys,
	})
}

func GetRecentHistoryData(c *gin.Context) {
	historys, err := service.GetRecentHistoryData()
	if err != nil {
		msg := fmt.Sprintf("api-GetRecentHistoryData-获取历史战绩错误")
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  msg,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": historys,
	})
}

func InsertHistoryData(c *gin.Context) {
	score, _ := strconv.Atoi(c.PostForm("length"))
	modeName := c.PostForm("modeName")
	fmt.Println("当前模式:", modeName)
	if modeName == "" {
		msg := fmt.Sprintf("api-InsertScore-获取mode字段失败")
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  msg,
		})
		return
	}

	err := service.InsertHistoryData(score, modeName)
	if err != nil {
		msg := fmt.Sprintf("api-插入成绩错误")
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  msg,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
	})
}
