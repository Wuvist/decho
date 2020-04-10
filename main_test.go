package main

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	config, err := loadTomlConf()
	if err != nil {
		t.Error(err)
	}

	if config.Address != ":1323" {
		t.Errorf("Invalid address: %s", config.Address)
	}
}
