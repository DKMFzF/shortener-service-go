package main

import (
	"fmt"
	"shortener/internal/configs/apps/clientConfig"
	"strings"

	"github.com/go-resty/resty/v2"
)

func main() {
	endpoint := clientConfig.New()

	longUrl := strings.TrimSuffix(endpoint.Url, "\n")
	fmt.Println("url req:", longUrl)

	client := resty.New()

	// TODO: добавить кастом под каждый эндпоинт
	resp, err := client.R().
		SetHeader("Content-Type", "text/plain; charset=utf-8").
		SetFormData(map[string]string{
			"url": longUrl,
		}).
		Get(endpoint.Url)

	if err != nil {
		panic(err)
	}

	fmt.Println("status-code:", resp.StatusCode())
	fmt.Println("response:")
	fmt.Println(resp.String())
}
