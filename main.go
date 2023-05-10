package main

import (
	"github.com/alecthomas/kong"
	"github.com/eswulei/chatgpt-web/bootstarp"
	"github.com/eswulei/chatgpt-web/config"
)

func main() {
	kong.Parse(&config.CLI)
	bootstrap.StartWebServer()
}
