package main

import "github.com/Drelf2018/webhook"

func main() {
	webhook.Run(&webhook.Config{
		Path:           webhook.Path{Root: "/home/app"},
		Administrators: []string{"188888131"},
	})
}
