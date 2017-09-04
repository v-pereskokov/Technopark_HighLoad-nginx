package utils

import (
	"encoding/json"
	"fmt"
	"github.com/vladpereskokov/Technopark_HighLoad-nginx/configs"
	"io/ioutil"
)

func FromFile(filename string) (*configs.Config, error) {
	file, err := ioutil.ReadFile(filename)

	if err != nil {
		return &configs.Config{}, fmt.Errorf("can not open file: %v", err)
	}

	return fromReader(file)
}

func fromReader(r []byte) (*configs.Config, error) {
	config := new(configs.Config)
	err := json.Unmarshal(r, &config)

	if err != nil {
		return config, fmt.Errorf("can not parse config: %v", err)
	}

	return config, nil
}
