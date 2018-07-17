package application

import (
	"Go-KV/common"
	"github.com/patrickmn/go-cache"
)

type Application1 struct {
	Cache	*common.KVMemOnly
}

type Application1OnMisser struct {
	// todo mysql

}

func (onmisser *Application1OnMisser) OnMiss(key string, object interface{}) (result interface{}, err error) {
	return
}

func NewApplication1() (Application1) {
	cache := cache.Cache{

	}
	application1 := Application1{
		Cache:	&common.KVMemOnly{
			Cache: &cache,
			OnMisser:	&Application1OnMisser{},
		},
	}
	return application1
}
