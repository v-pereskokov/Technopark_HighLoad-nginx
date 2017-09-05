package utils

import (
	"os"
)

func Pwd() string {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return pwd
}
