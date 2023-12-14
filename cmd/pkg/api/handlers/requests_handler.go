package handlers

import (
	"example/snap_roles/configs"
	"io"
	"log"
	"net/http"
)

type HandlerInterface interface {
	GetApi(url string) (*http.Response, error)
	PostApi(url string, contentType string, body io.Reader) (*http.Response, error)
}

type HandlerStruct struct {
	configImpl configs.ConfigInterface
}

func (handler HandlerStruct) GetApi(url string) (*http.Response, error) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("GET API error %s. url: %s", err, url)
	}
	return response, nil

}

func (handler HandlerStruct) PostApi(url string, contentType string, body io.Reader) (*http.Response, error) {
	response, err := http.Post(url, contentType, body)
	if err != nil {
		log.Fatalf("POST API error %s. url: %s", err, url)
	}
	return response, nil
}
