package util

import (
	"io/ioutil"
	"os"
)

func CreateFileIfNotExist(filePath string) error {
	exist, err := FileExist(filePath)
	if err != nil {
		return err
	}

	if exist {
		return nil
	}

	f, err := os.Create(filePath)
	defer f.Close()

	if err != nil {
		return err
	}
	return nil
}

func FileExist(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func ClearAndWrite(filePath string, content string) error {
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0600)
	defer f.Close()

	if err != nil {
		return err
	}

	_, err1 := f.WriteString(content)
	if err1 != nil {
		return err1
	}

	return nil
}

func ReadFileAll(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)

	if err != nil {
		return "", err
	}
	return string(content), nil
}
