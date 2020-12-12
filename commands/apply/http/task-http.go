package http

import (
	"fmt"
	"github.com/sumittokkar/arrow/model"
	"io/ioutil"
	"net/http"
)

func ProcessTaskTypeHttp(taskType *model.TaskType) error {
	h := taskType.Http

	if h == nil {
		return nil
	}

	resp, err := execute(h)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)

		fmt.Println(string(bodyBytes))
	}

	return nil
}
