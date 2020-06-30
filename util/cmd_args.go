package util

import "flag"

func BindIntCmdParam(name string) *int64 {
	return flag.Int64(name, 0, "")
}

func BindStringCmdParam(name string) *string {
	return flag.String(name, "", "")
}
