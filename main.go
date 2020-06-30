package main

import (
	"api-mock/handler"
	"api-mock/mock_request"
	"api-mock/util"
	"flag"
	"github.com/gin-gonic/gin"
	"sync"
)

func main() {
	cPort := util.BindStringCmdParam("p")
	flag.Parse()
	port := ""
	if *cPort == "" {
		util.InfoLog("-p Not defined , use default val 8090", port)
		port = "8090"
	} else {
		port = *cPort
		util.InfoLog("Use port : "+port, "")
	}

	InitMock()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		InitGin(port)
		wg.Done()
	}()

	util.InfoLog("Start Finish...", "")
	wg.Wait()
	util.InfoLog("App Stopped...", "")
}

func InitMock() {
	mappingFile := "./data/mockReqRespMappings"
	err := mock_request.InitMockReqRespMappings(mappingFile)
	if err != nil {
		util.ErrorLog("InitMockReqRespMappings error :"+err.Error(), mappingFile)
	} else {
		util.InfoLog("InitMockReqRespMappings success ", mappingFile)
	}
}

func InitGin(port string) {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	engine.Use(handler.BeforeFilter)

	engine.GET("/mapping/list", handler.ListMappingHandler)
	engine.POST("/mapping/save", handler.SaveMappingHandler)
	engine.POST("/mapping/del", handler.DelMappingHandler)

	err := engine.Run(":" + port)
	if err != nil {
		util.ErrorLog("Start HttpServer Fail :"+err.Error(), "")
	}
}
