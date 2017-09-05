package configs

import (
	"strings"
	"testing"
)

const file_1 = "config_test_1.json"
const file_2 = "config_test_2.json"

func TestFromFileFirst(t *testing.T) {
	config, err := FromFile(file_1)
	if err != nil {
		t.Errorf("%v\n", err.Error())
	}

	if config.GetNetwork() != "tcp" {
		t.Errorf("network don't match! %s", config.GetNetwork())
	}

	if config.GetProtocol() != "http" {
		t.Errorf("protocol don't match! %s", config.GetProtocol())
	}

	if config.GetHost() != "localhost" {
		t.Errorf("host don't match! %s", config.GetHost())
	}

	if config.GetPort() != "2007" {
		t.Errorf("port don't match! %s", config.GetPort())
	}

	if config.GetDir() != "/Users/vladislavpereskokov/Desktop/" {
		t.Errorf("dir don't match! %s", config.GetPort())
	}
}

func TestFromFileSecond(t *testing.T) {
	config, err := FromFile(file_2)
	if err != nil {
		t.Errorf("%v\n", err.Error())
	}

	if config.GetNetwork() != "tcp" {
		t.Errorf("network don't match! %s", config.GetNetwork())
	}

	if config.GetProtocol() != "http" {
		t.Errorf("protocol don't match! %s", config.GetProtocol())
	}

	if config.GetHost() != "localhost" {
		t.Errorf("host don't match! %s", config.GetHost())
	}

	if config.GetPort() != "2007" {
		t.Errorf("port don't match! %s", config.GetPort())
	}

	if !strings.Contains(config.GetDir(), "github.com/vladpereskokov/Technopark_HighLoad-nginx/") {
		t.Errorf("pwd dir don't match! %s", config.GetPort())
	}
}
