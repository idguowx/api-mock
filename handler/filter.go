package handler

import (
	"api-mock/mock_request"
	"api-mock/util"
	"github.com/gin-gonic/gin"
)

func BeforeFilter(context *gin.Context) {
	if util.RegMatch("^/mapping/save", context.Request.RequestURI) ||
		util.RegMatch("^/mapping/del", context.Request.RequestURI) ||
		util.RegMatch("^/mapping/list", context.Request.RequestURI) {
		context.Next()
		return
	}
	if context.Request.RequestURI == "/" {
		context.String(200, "ok")
		return
	}

	request := mock_request.MakeMockRequest()
	request.FillMockRequest(context)

	response, err := request.MatchResponse()
	util.InfoLog("MatchRequest", map[string]interface{}{
		"request":  request,
		"response": response,
	})
	if err != nil {
		context.String(200, err.Error())
		return
	}

	context.String(200, response)
	return
}
