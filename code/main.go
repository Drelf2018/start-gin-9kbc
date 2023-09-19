package main

import (
	"github.com/Drelf2018/webhook"
	"github.com/Drelf2018/webhook/configs"
)

func main() {
	webhook.Release(&webhook.Config{
		Path:           configs.Path{Root: "/home/app"},
		Administrators: []string{"188888131"},
	})
}
