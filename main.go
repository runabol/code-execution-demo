package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/runabol/code-execution-demo/handler"
	"github.com/runabol/tork/cli"
	"github.com/runabol/tork/conf"
)

func main() {
	if err := conf.LoadConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	app := cli.New()

	app.RegisterEndpoint(http.MethodPost, "/execute", handler.Handler)

	if err := app.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
