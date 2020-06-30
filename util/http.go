package util

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type HttpResponse struct {
	Success      bool
	ErrorCode    int
	ErrorMessage string
	Payload      interface{}
}

func (r *HttpResponse) SuccessResponse(context *gin.Context) {
	r.Success = true
	jsonBytes, _ := json.Marshal(r)
	context.String(200, string(jsonBytes))
}

func (r *HttpResponse) ErrorResponse(context *gin.Context) {
	r.Success = false
	jsonBytes, _ := json.Marshal(r)
	context.String(200, string(jsonBytes))
}
