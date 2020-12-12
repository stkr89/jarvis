package commands

import (
	"bytes"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gookit/color"
	"github.com/stkr89/jarvis/model"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func Verify() func(c *cli.Context) error {
	return func(c *cli.Context) error {
		if c.Args().Len() != 1 {
			color.Red.Println("path to config file is required")
			return nil
		}

		_, err := ProcessVerify(c.Args().First())
		if err != nil {
			color.Red.Println(err.Error())
			return nil
		}

		color.Green.Println("Looks good!")

		return nil
	}
}

func ProcessVerify(path string) (*model.Tasks, error) {
	fileBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var tasks model.Tasks
	err = yaml.Unmarshal(fileBytes, &tasks)
	if err != nil {
		return nil, err
	}

	err = validator.New().Struct(tasks)
	if err != nil {
		errStr := formatValidationError(err)
		if errStr != "" {
			return nil, errors.New(errStr)
		}

		return nil, nil
	}

	return &tasks, nil
}

func formatValidationError(err error) string {
	var buffer bytes.Buffer

	if _, ok := err.(*validator.InvalidValidationError); ok {
		return buffer.String()
	}

	for _, err := range err.(validator.ValidationErrors) {
		buffer.WriteString(err.Namespace() + ":" + err.ActualTag() + ";")
	}

	return buffer.String()
}
