package configs

import (
	"strings"
	"testing"
)

const (
	FILE_1 = "../../test_configs/config_test_1.json"
	FILE_2 = "../../test_configs/config_test_2.json"
)

func TestFromFileFirst(t *testing.T) {
	config, err := FromFile(FILE_1)
	if err != nil {
		t.Errorf("%v\n", err.Error())
	}

	serverConfig := config.GetServer()
	dirConfig := config.GetDir()

	if serverConfig.GetNetwork() != "tcp" {
		t.Errorf("network don't match! %s", serverConfig.GetNetwork())
	}

	if serverConfig.GetProtocol() != "http" {
		t.Errorf("protocol don't match! %s", serverConfig.GetProtocol())
	}

	if serverConfig.GetHost() != "localhost" {
		t.Errorf("host don't match! %s", serverConfig.GetHost())
	}

	if serverConfig.GetPort() != "2007" {
		t.Errorf("port don't match! %s", serverConfig.GetPort())
	}

	if dirConfig.GetDir() != "/usr/topDir/" {
		t.Errorf("dir don't match! %s", dirConfig.GetDir())
	}
}

func TestFromFileSecond(t *testing.T) {
	config, err := FromFile(FILE_2)
	if err != nil {
		t.Errorf("%v\n", err.Error())
	}

	serverConfig := config.GetServer()
	dirConfig := config.GetDir()

	if serverConfig.GetNetwork() != "tcp" {
		t.Errorf("network don't match! %s", serverConfig.GetNetwork())
	}

	if serverConfig.GetProtocol() != "http" {
		t.Errorf("protocol don't match! %s", serverConfig.GetProtocol())
	}

	if serverConfig.GetHost() != "localhost" {
		t.Errorf("host don't match! %s", serverConfig.GetHost())
	}

	if serverConfig.GetPort() != "2007" {
		t.Errorf("port don't match! %s", serverConfig.GetPort())
	}

	if !strings.Contains(dirConfig.GetDir(), "github.com/vladpereskokov/Technopark_HighLoad-nginx/") {
		t.Errorf("pwd dir don't match! %s", dirConfig.GetDir())
	}
}
