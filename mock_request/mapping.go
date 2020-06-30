package mock_request

import (
	"api-mock/util"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

//for example
//[
//	"appName":[
//		[
//			"requestPattern":[
//			],
//			"requestPattern":[
//			],
//		]
//	]
//]
var mockReqRespMappings = map[string]map[string]*ReqRespMapping{}
var serializeFile = "../data/mockReqRespMappings"

type ReqRespMapping struct {
	CaseName       string
	RequestPattern *RequestPattern
	ResponseStr    string
}

func InitMockReqRespMappings(serializeFile string) error {
	SetSerializeFile(serializeFile)
	return buildMappingsFromSerializeFile()
}

func (m *ReqRespMapping) isRequestMatched(req *Request) bool {
	pattern := m.RequestPattern

	if pattern.Method.Ignore == util.NUM_FALSE &&
		strings.ToLower(strings.TrimSpace(pattern.Method.Value)) !=
			strings.ToLower(strings.TrimSpace(req.Method)) {
		return false
	}

	if pattern.Method.Ignore == util.NUM_FALSE && !util.RegMatch(pattern.ContentTypeReg.Value, req.ContentType) {
		fmt.Println(pattern.ContentTypeReg.Value)
		fmt.Println(req.ContentType)
		return false
	}

	if pattern.Method.Ignore == util.NUM_FALSE && !util.RegMatch(pattern.BodyReg.Value, req.Body) {
		return false
	}

	return true
}

func GetAppReqMappings(appName string) (map[string]*ReqRespMapping, error) {
	if _, ok := mockReqRespMappings[appName]; !ok {
		return nil, errors.New("app not found")
	}

	return mockReqRespMappings[appName], nil
}

func GetAllAppReqMappings() map[string]map[string]*ReqRespMapping {
	fmt.Println(len(mockReqRespMappings))
	return mockReqRespMappings
}

func AddAppReqMappings(appName string, reqRespMapping *ReqRespMapping) map[string]*ReqRespMapping {
	if _, ok := mockReqRespMappings[appName]; !ok {
		mockReqRespMappings[appName] = make(map[string]*ReqRespMapping)
	}

	mockReqRespMappings[appName][reqRespMapping.CaseName] = reqRespMapping
	saveMappingsToSerializeFile()
	return mockReqRespMappings[appName]
}

func DelAppReqMappings(appName string, caseName string) (map[string]*ReqRespMapping, error) {
	if _, ok := mockReqRespMappings[appName]; !ok {
		mockReqRespMappings[appName] = make(map[string]*ReqRespMapping, 0)
	}

	delete(mockReqRespMappings[appName], caseName)
	saveMappingsToSerializeFile()
	return mockReqRespMappings[appName], nil
}

func SetSerializeFile(file string) {
	serializeFile = file
}

func saveMappingsToSerializeFile() error {
	err := util.CreateFileIfNotExist(serializeFile)
	if err != nil {
		return err
	}

	jsonBytes, _ := json.Marshal(mockReqRespMappings)
	return util.ClearAndWrite(serializeFile, string(jsonBytes))
}

func buildMappingsFromSerializeFile() error {
	jsonStr, err := util.ReadFileAll(serializeFile)
	if err != nil {
		return err
	}

	if jsonStr == "" {
		return nil
	}

	err = json.Unmarshal([]byte(jsonStr), &mockReqRespMappings)
	if err != nil {
		return err
	}

	return nil
}
