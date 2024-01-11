package cronfunc

import (
	"github.com/mileusna/crontab"
)

// https://crontab.guru/

func Add() *cronFunc {
	return &cronFunc{
		Crontab:    crontab.New(),
		registered: map[string]string{},
	}
}
