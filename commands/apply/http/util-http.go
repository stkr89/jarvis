package http

import (
	"encoding/base64"
	"github.com/sumittokkar/arrow/model"
	"net/http"
	"time"
)

func execute(h *model.TaskTypeHttp) (*http.Response, error) {
	req, _ := http.NewRequest(h.Method, h.Url, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	if h.Auth != nil {
		if h.Auth.Basic != nil {
			req.Header.Add("Authorization", "Basic "+basicAuth(h.Auth.Basic.Username, h.Auth.Basic.Password))
		}
	}

	client := &http.Client{Timeout: time.Second * time.Duration(h.Timeout)}

	return client.Do(req)
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
