package main

import "github.com/Drelf2018/webhook"

func main() {
	webhook.Release(&webhook.Config{
		Resource:       webhook.Resource{Path: "/home/app"},
		Administrators: []string{"188888131"},
	})
}
