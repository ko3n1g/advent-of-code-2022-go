package tools

import (
	"os"
	"strings"
)

func GetData(fileName string) []string {
	currPath, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	file, err := os.ReadFile(currPath + "/data/" + fileName)

	if err != nil {
		panic(err)
	}

	return strings.Split(string(file), "\n")
}
