package utils

import "testing"

const path = "config_test.json"

func TestFromFile(t *testing.T) {
	config, err := FromFile(path)
	if err != nil {
		t.Errorf("%v\n", err.Error())
	}

	if config.GetPort() != "3000" {
		t.Error("Don't match!")
	}
}
