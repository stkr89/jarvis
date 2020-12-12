package http

import (
	"github.com/stkr89/jarvis/model"
)

func ProcessTaskTypeHttp(taskType *model.TaskType) (string, error) {
	h := taskType.Http

	if h == nil {
		return "", nil
	}

	resp, err := execute(h)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return resp.Status, nil
}
