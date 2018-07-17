package common

import (
	"github.com/patrickmn/go-cache"
	"Go-KV/util"
	"time"
)

type KVMemOnly struct {
	Cache *cache.Cache
	OnMisser MemOnlyOnMisser
	// 其他需要的参数
}

type MemOnlyOnMisser interface {
	OnMiss(key string, object interface{}) (result interface {}, err error)
}

func (kv *KVMemOnly) Get(key string, object interface{}) (result interface{}, err error) {
	log := *(util.Get())
	// hit
	if value, found := kv.Cache.Get(key); found {
		result = value
		return
	}
	// not hit
	value, err := kv.OnMisser.OnMiss(key, object)
	if err != nil {
		log.Error(err)
		value = nil
	}
	kv.Cache.Set(key, value, time.Duration(0))
	result = value
	return
}
