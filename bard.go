package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

const BardAPI = "https://api.bardapi.dev/chat"

type Response struct {
	Error string `json:"error"`
	Output string 		`json:"output"`
}

type ErrorResponse struct {
	Data string `json:"data"`
}

type BardInterface interface {
	GetAnswers(question string)	(string, error)
}

type BardModel struct {
	key string
}

func Bard(bearerKey string) BardInterface {
	return &BardModel{
		key: bearerKey,
	}
}

func (b *BardModel) GetAnswers(question string) (string, error) {
	questionBody := map[string]string{
		"input": question,
	}
	values, _ := json.Marshal(questionBody)
	
	req, err := http.NewRequest("POST", BardAPI, bytes.NewBuffer(values))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("Authorization", "Bearer " + b.key)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var res Response
	json.NewDecoder(resp.Body).Decode(&res)
	if res.Error != "" {
		return "", errors.New(res.Error) 
	}
	return res.Output, nil
}