package api

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(Cors())

	history := router.Group("api/history")
	{
		history.GET("/getBestHistoryDataByMode/:modeName", GetBestHistoryDataByMode)
		history.GET("/getBestHistoryData", GetBestHistoryData)
		history.GET("/getRecentHistoryData", GetRecentHistoryData)
		history.POST("/insertHistoryData", InsertHistoryData)
	}
	return router
}
