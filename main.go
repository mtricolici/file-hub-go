package main

import (
	"fmt"
	"os"

	"github.com/mtricolici/file-hub-go/config"
)

func main() {
	// testing
	os.Setenv("CONFIG_PATH", "~/Projects/file-hub-go/config-sample.yml")
	config.Get().LoadConfig()

	listen := config.Get().ListenInterface()
	port := config.Get().ListenPort()
	fmt.Printf("listen on '%s:%d'\n", listen, port)
}
