package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/runabol/code-execution-demo/handler"
	"github.com/runabol/tork/cli"
	"github.com/runabol/tork/pkg/conf"
	"github.com/runabol/tork/pkg/engine"
)

func main() {
	if err := conf.LoadConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	app := cli.New()

	app.ConfigureEngine(func(eng *engine.Engine) error {
		eng.RegisterEndpoint(http.MethodPost, "/execute", handler.Handler)
		return nil
	})

	if err := app.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
