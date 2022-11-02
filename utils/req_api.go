package utils

import (
	"bytes"
	"log"
	"net/http"
	"time"
)

func SendGetAsync(url string, rc chan *http.Response) error {
	req, err1 := http.NewRequest(http.MethodGet, url, bytes.NewBuffer([]byte("")))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer ")

	if err1 != nil {
		log.Println(err1)
		return err1
	}

	client := &http.Client{}
	client.Timeout = time.Second * 3
	response, err := client.Do(req)

	if err == nil {
		rc <- response
	}
	return err
}

func SendPostAsync(url string, body []byte, rc chan *http.Response) error {
	var jsonStr = []byte(body)
	req, err1 := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	if err1 != nil {
		log.Println(err1)
		return err1
	}

	client := &http.Client{}
	client.Timeout = time.Second * 3
	response, err := client.Do(req)

	if err == nil {
		rc <- response
	}
	return err
}

func SendLineReply(url string, body []byte, rc chan *http.Response, token string) error {
	var jsonStr = []byte(body)
	req, err1 := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	if err1 != nil {
		log.Println(err1)
		return err1
	}

	client := &http.Client{}
	client.Timeout = time.Second * 3
	response, err := client.Do(req)

	if err == nil {
		rc <- response
	}
	return err
}

func SendPutAsync(url string, body string, rc chan *http.Response) error {
	var jsonStr = []byte(body)
	req, err1 := http.NewRequest(http.MethodPut, url, bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	if err1 != nil {
		log.Println(err1)
		return err1
	}

	client := &http.Client{}
	response, err := client.Do(req)

	if err == nil {
		rc <- response
	}
	return err
}

func SendDeleteAsync(url string, body string, rc chan *http.Response) error {
	var jsonStr = []byte(body)
	req, err1 := http.NewRequest(http.MethodDelete, url, bytes.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	if err1 != nil {
		log.Println(err1)
		return err1
	}

	client := &http.Client{}
	response, err := client.Do(req)

	if err == nil {
		rc <- response
	}
	return err
}
