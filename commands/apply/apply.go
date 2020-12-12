package apply

import (
	"github.com/gookit/color"
	"github.com/stkr89/jarvis/commands"
	"github.com/stkr89/jarvis/commands/apply/http"
	"github.com/stkr89/jarvis/model"
	"github.com/urfave/cli/v2"
)

func Apply() func(c *cli.Context) error {
	return func(c *cli.Context) error {
		if c.Args().Len() != 1 {
			color.Red.Println("path to config file is required")
			return nil
		}

		tasks, err := commands.ProcessVerify(c.Args().First())
		if err != nil {
			color.Red.Println(err.Error())
			return nil
		}

		err = processApply(tasks)
		if err != nil {
			color.Red.Println(err.Error())
			return nil
		}

		return nil
	}
}

func processApply(tasks *model.Tasks) error {
	processors := getProcessors()

	for _, t := range tasks.Tasks {
		color.Green.Printf("executing task: %s\n", t.Name)

		for _, processor := range processors {
			response, err := processor(t.TaskType)
			if err != nil {
				return err
			}

			color.Green.Println(response)
		}
	}

	return nil
}

func getProcessors() []func(taskType *model.TaskType) (string, error) {
	return []func(taskType *model.TaskType) (string, error){
		http.ProcessTaskTypeHttp,
	}
}
