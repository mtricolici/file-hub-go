package main

import (
	"os"

	"github.com/mtricolici/file-hub-go/config"
	"github.com/mtricolici/file-hub-go/web"
)

func main() {
	// testing
	os.Setenv("CONFIG_PATH", "~/Projects/file-hub-go/config-sample.yml")
	config.Get().LoadConfig()

	web.InitAndStart()
}
