package handler

import (
	"api-mock/mock_request"
	"api-mock/util"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func ListMappingHandler(context *gin.Context) {
	appName := context.Query("testAppName")
	r := &util.HttpResponse{}
	if appName == "" {
		r.Payload = mock_request.GetAllAppReqMappings()
	} else {
		r.Payload, _ = mock_request.GetAppReqMappings(appName)
	}

	r.SuccessResponse(context)
}

func SaveMappingHandler(context *gin.Context) {
	appName := context.PostForm("testAppName")
	caseName := context.PostForm("caseName")
	requestPattern := context.PostForm("requestPattern")
	responseStr := context.PostForm("responseStr")

	r := &util.HttpResponse{}
	if appName == "" || caseName == "" || requestPattern == "" || responseStr == "" {
		r.ErrorMessage = "param error"
		r.ErrorResponse(context)
		return
	}
	pattern := &mock_request.RequestPattern{}
	err := json.Unmarshal([]byte(requestPattern), pattern)
	if err != nil {
		r.ErrorMessage = "requestPattern format error"
		r.ErrorResponse(context)
		return
	}
	mock_request.AddAppReqMappings(appName, &mock_request.ReqRespMapping{
		CaseName:       caseName,
		RequestPattern: pattern,
		ResponseStr:    responseStr,
	})
	r.SuccessResponse(context)
}

func DelMappingHandler(context *gin.Context) {
	appName := context.PostForm("testAppName")
	caseName := context.PostForm("caseName")

	r := &util.HttpResponse{}
	if appName == "" || caseName == "" {
		r.ErrorMessage = "param error"
		r.ErrorResponse(context)
		return
	}

	mock_request.DelAppReqMappings(appName, caseName)
	r.SuccessResponse(context)
}
