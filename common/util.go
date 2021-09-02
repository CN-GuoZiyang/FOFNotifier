package common

import (
	"FOFNotifier/model"
	"fmt"
	"sync"
)

var (
	once   sync.Once
	Config = &model.GlobalConfig{}
)

func GetGlobalConfig() *model.GlobalConfig {
	once.Do(func() {
		combinationConfig, err := ReadCombinationConfig()
		if err != nil {
			fmt.Printf("Read combination config error, err=%v\n", err)
		}
		Config.Combinations = combinationConfig
	})
	return Config
}
