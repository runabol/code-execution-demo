package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/runabol/tork-demo-codexec/handler"
	"github.com/runabol/tork/bootstrap"
	"github.com/runabol/tork/cli"
	"github.com/runabol/tork/conf"
)

func main() {
	if err := conf.LoadConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bootstrap.RegisterEndpoint(http.MethodPost, "/exec", handler.Handler)

	if err := cli.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
