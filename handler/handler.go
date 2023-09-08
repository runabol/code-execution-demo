package handler

import (
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/runabol/tork"
	"github.com/runabol/tork/input"
	"github.com/runabol/tork/middleware"
)

type ExecRequest struct {
	Code     string `json:"code"`
	Language string `json:"language"`
}

func Handler(c middleware.Context) error {
	er := ExecRequest{}

	if err := c.Bind(&er); err != nil {
		return err
	}

	log.Debug().Msgf("%s", er.Code)

	task, err := buildTask(er)
	if err != nil {
		return err
	}

	result := make(chan string)

	listener := func(j *tork.Job) {
		if j.State == tork.JobStateCompleted {
			result <- j.Execution[0].Result
		} else {
			result <- j.Execution[0].Error
		}
	}

	input := &input.Job{
		Name:  "code execution",
		Tasks: []input.Task{task},
	}

	job, err := c.SubmitJob(input, listener)

	if err != nil {
		c.Error(http.StatusBadRequest, errors.Wrapf(err, "error executing code"))
		return nil
	}

	log.Debug().Msgf("job %s submitted", job.ID)

	return c.String(http.StatusOK, <-result)
}

func buildTask(er ExecRequest) (input.Task, error) {
	var image string
	var script string
	var filename string

	switch strings.TrimSpace(er.Language) {
	case "":
		return input.Task{}, errors.Errorf("require: language")
	case "python3":
		image = "python:3"
		filename = "script.py"
		script = "python script.py > $TORK_OUTPUT"
	case "golang":
		image = "golang:1.19"
		filename = "main.go"
		script = "go run main.go > $TORK_OUTPUT"
	default:
		return input.Task{}, errors.Errorf("unknown language: %s", er.Language)
	}

	return input.Task{
		Name:  "execute code",
		Image: image,
		Run:   script,
		Files: map[string]string{
			filename: er.Code,
		},
	}, nil
}