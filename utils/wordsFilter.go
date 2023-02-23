package utils

import (
	"dousheng/config"
	"github.com/importcjj/sensitive"
	"log"
)

var Filter *sensitive.Filter

func FilterInit() {
	Filter = sensitive.New()
	err := Filter.LoadWordDict(config.DICT)
	if err != nil {
		log.Panicln(err)
	}
}
