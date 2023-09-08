package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/runabol/tork"
	"github.com/runabol/tork/bootstrap"
	"github.com/runabol/tork/cli"
	"github.com/runabol/tork/conf"
	"github.com/runabol/tork/input"
	"github.com/runabol/tork/middleware"
)

func main() {
	if err := conf.LoadConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bootstrap.RegisterEndpoint(http.MethodPost, "/exec", handleExec)

	if err := cli.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func handleExec(c middleware.Context) error {
	// bind request to a struct

	type ExecRequest struct {
		Code     string `json:"code"`
		Language string `json:"language"`
	}

	er := ExecRequest{}

	if err := c.Bind(&er); err != nil {
		return err
	}

	var image string
	switch er.Language {
	case "python3":
		image = "python:3"
	case "":
		c.Error(http.StatusBadRequest, errors.Errorf("require: language"))
		return nil
	default:
		c.Error(http.StatusBadRequest, errors.Errorf("unknown language: %s", er.Language))
		return nil
	}

	if er.Code == "" {
		c.Error(http.StatusBadRequest, errors.Errorf("require: code"))
		return nil
	}

	task := input.Task{
		Name:  "execute code",
		Image: image,
		Run:   "python script.py > $TORK_OUTPUT",
		Files: map[string]string{
			"script.py": er.Code,
		},
	}

	result := make(chan string)

	listener := func(j *tork.Job) {
		if j.State == tork.JobStateCompleted {
			result <- j.Execution[0].Result
		} else {
			result <- j.Execution[0].Error
		}
	}

	job, err := c.SubmitJob(&input.Job{
		Name:  "code execution",
		Tasks: []input.Task{task},
	}, listener)

	if err != nil {
		c.Error(http.StatusBadRequest, errors.Wrapf(err, "error executing code"))
		return nil
	}

	log.Debug().Msgf("job %s submitted", job.ID)

	// send execution output back to client

	return c.String(http.StatusOK, <-result)
}
