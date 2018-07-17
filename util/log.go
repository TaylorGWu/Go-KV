package util

import (
	"Go-KV/config"
	log "github.com/cihub/seelog"
)

var logger log.LoggerInterface

func New() {
	goKVConfig := config.Get()

	var err error
	logger, err = log.LoggerFromConfigAsFile(goKVConfig.LogConfigPath)
	if err != nil {
		panic("init seelog fail")
	}

	log.ReplaceLogger(logger)
}

func Del() {
	logger.Flush()
}

func Get() *log.LoggerInterface {
	return &logger
}


