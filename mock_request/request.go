package mock_request

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type Request struct {
	AppName     string
	Method      string
	Uri         string
	ContentType string
	Body        string
}

func MakeMockRequest() *Request {
	return &Request{
		AppName:     "default",
		Method:      "",
		Uri:         "",
		ContentType: "",
		Body:        "",
	}
}

//fill request
func (r *Request) FillMockRequest(context *gin.Context) *Request {
	r.ContentType = context.ContentType()
	r.Uri = context.Request.RequestURI
	r.Method = context.Request.Method

	appName := context.Query("testAppName")
	if appName != "" {
		r.AppName = appName
	}

	bodyBytes, _ := ioutil.ReadAll(context.Request.Body)
	r.Body = string(bodyBytes)

	return r
}

func (r *Request) MatchResponse() (string, error) {
	mappings, err := GetAppReqMappings(r.AppName)
	if err != nil {
		return "", err
	}

	for _, m := range mappings {
		if m.isRequestMatched(r) {
			return m.ResponseStr, nil
		}
	}
	return "", errors.New("not match")
}

func (r *Request) ToString() string {
	return fmt.Sprintf("Method=[%s] ,Uri=[%s] ,ContentType=[%s] ,Body=[%s]",
		r.Method, r.Uri, r.ContentType, r.Body)
}
