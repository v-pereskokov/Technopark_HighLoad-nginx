package utils

import "testing"

const path = "config_test.json"

func TestFromFile(t *testing.T) {
	config, err := FromFile(path)
	if err != nil {
		t.Errorf("%v\n", err.Error())
	}

	if config.GetNetwork() != "tcp" {
		t.Errorf("host don't match! %s", config.GetNetwork())
	}

	if config.GetProtocol() != "http" {
		t.Errorf("host don't match! %s", config.GetProtocol())
	}

	if config.GetHost() != "localhost" {
		t.Errorf("host don't match! %s", config.GetHost())
	}

	if config.GetPort() != "2007" {
		t.Errorf("port don't match! %s", config.GetPort())
	}
}
