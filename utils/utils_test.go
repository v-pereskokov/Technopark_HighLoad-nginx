package utils

import "testing"

const path = "config_test.json"

func TestFromFile(t *testing.T) {
	config, err := FromFile(path)
	if err != nil {
		t.Errorf("%v\n", err.Error())
	}

	if config.GetPort() != "2007" {
		t.Errorf("Don't match! %s", config.GetPort())
	}
}
