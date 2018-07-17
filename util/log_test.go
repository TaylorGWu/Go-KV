package util

import (
	"testing"
	"Go-KV/config"
)

func TestLog(t *testing.T) {
	config.ConfigGoKV()
	New()
	defer Del()
	log := Get()
	(*log).Debugf("hello")
}
