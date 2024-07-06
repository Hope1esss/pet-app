package service

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/json-iterator/go"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func getAccessToken() string {
	if err := godotenv.Load("./internal/config/gigachat.env"); err != nil {
		log.Print("No .env file found")
	}

	url := "https://ngw.devices.sberbank.ru:9443/api/v2/oauth"
	method := "POST"

	payload := strings.NewReader("scope=GIGACHAT_API_PERS")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("RqUID", os.Getenv("RQUID"))
	req.Header.Add("Authorization", os.Getenv("OAUTH_TOKEN"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	var response map[string]interface{}
	err = jsoniter.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return response["access_token"].(string)

}
func prepareInputMessage(inputMessage string) []byte {
	prompt := "Ты профессиональный зоолог. Твоя задача – подобрать пользователю домашнее животное, максимально удовлетворяющее его пожеланиям. Ответ должен быть такого формата: \n1) Ответ состоит МИНИМУМ из 5 пунктов. Чем больше - тем лучше;\n2) К каждому варианту должно прилагаться описание этого варианта, которое рассказывает как о положительных сторонах, так и об отрицательных. Объём – от 20 до 50 слов. \n"
	type Message struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	}
	type Body struct {
		Model          string    `json:"model"`
		Messages       []Message `json:"messages"`
		N              int       `json:"n"`
		Stream         bool      `json:"stream"`
		UpdateInterval int       `json:"update_interval"`
	}
	data := Body{
		Model: "GigaChat",
		Messages: []Message{
			{Role: "system", Content: prompt},
			{Role: "user", Content: inputMessage},
		},
		N:              1,
		Stream:         false,
		UpdateInterval: 0,
	}
	jsonBody, err := jsoniter.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return jsonBody
}

func prepareOutputMessage(body []byte) string {

	type Message struct {
		Content string `json:"content"`
		Role    string `json:"role"`
	}
	type Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	}

	type Choice struct {
		Message      Message `json:"message"`
		Index        int     `json:"index"`
		FinishReason string  `json:"finish_reason"`
	}

	type Response struct {
		Choices []Choice `json:"choices"`
		Created int      `json:"created"`
		Model   string   `json:"model"`
		Object  string   `json:"object"`
		Usage   Usage    `json:"usage"`
	}

	var response Response
	err := jsoniter.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return response.Choices[0].Message.Content
}
func messageToGigaChat(jsonBody []byte) []byte {
	url := "https://gigachat.devices.sberbank.ru/api/v1/chat/completions"
	method := "POST"
	payload := strings.NewReader(string(jsonBody))

	client := &http.Client{}

	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+getAccessToken())

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil

	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return body
}
func Giga(inputMessage string) string {
	return prepareOutputMessage(messageToGigaChat(prepareInputMessage(inputMessage)))
}
