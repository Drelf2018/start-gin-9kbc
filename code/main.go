package main

import "github.com/Drelf2018/webhook"

func main() {
	webhook.Run(&webhook.Config{
		Resource: webhook.Resource{Path: "/home/app"},
		Github:   webhook.Github{Repository: "gin.nana7mi.link"},
	})
}
