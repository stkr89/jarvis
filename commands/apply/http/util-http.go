package http

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/stkr89/jarvis/model"
	"net/http"
	"regexp"
	"time"
)

func execute(h *model.TaskTypeHttp) (*http.Response, error) {
	if h.Body == nil {
		h.Body = make(map[string]string)
	}

	jsonBytes, err := json.Marshal(h.Body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(h.Method, h.Url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	if h.Auth != nil {
		addAuthorizationHeaders(h, req)
	}

	client := &http.Client{Timeout: time.Second * time.Duration(h.Timeout)}

	return client.Do(req)
}

func addAuthorizationHeaders(h *model.TaskTypeHttp, req *http.Request) {
	regex := regexp.MustCompile("^(\\${)?(.*?)}?$")

	if h.Auth.Basic != nil {
		req.Header.Add("Authorization", "Basic "+basicAuth(h.Auth.Basic.Username, h.Auth.Basic.Password, regex))
	} else if h.Auth.BearerToken != nil {
		req.Header.Add("Authorization", "Bearer "+parseStringFromEnv(h.Auth.BearerToken.Token, regex))
	} else if h.Auth.Custom != nil {
		for k, v := range h.Auth.Custom {
			req.Header.Add(k, parseStringFromEnv(v, regex))
		}
	}
}

func basicAuth(username, password string, regex *regexp.Regexp) string {
	auth := parseStringFromEnv(username, regex) + ":" + parseStringFromEnv(password, regex)
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func parseStringFromEnv(str string, regex *regexp.Regexp) string {
	u := regex.FindAllStringSubmatch(str, -1)
	if len(u[0]) == 3 {
		str = u[0][2]
	} else {
		str = u[0][1]
	}
	return str
}
