package main

import (
	"os/exec"

	"github.com/Drelf2018/webhook"
)

func main() {
	exec.Command("git", "clone", "-b", "gh-pages", "https://github.com/Drelf2018/nana7mi.link.git").Run()
	webhook.Run(&webhook.Config{})
}
