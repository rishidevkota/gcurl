package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Request struct {
	Method  string                  `json:"method"`
	URL     string                  `json:"url"`
	Headers *map[string]string      `json:"headers"`
	Body    *map[string]interface{} `json:"body"`
}

func Do(request Request) error {
	client := &http.Client{}
	var req *http.Request
	var err error

	switch request.Method {
	case "GET":
		req, err = http.NewRequest(request.Method, request.URL, nil)
		if err != nil {
			return err
		}
	case "POST":
		b, err := json.Marshal(request.Body)
		if err != nil {
			return err
		}
		req, err = http.NewRequest(request.Method, request.URL, bytes.NewReader(b))
		if err != nil {
			return err
		}
	}

	if request.Headers != nil {
		for key, value := range *request.Headers {
			req.Header.Add(key, value)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))

	return nil
}
