package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/go-resty/resty/v2"
)

const (
	baseUrl = "http://localhost:8080/"
)

func main() {
	endpoint := "http://localhost:8080/"

	fmt.Println("Введите длинный URL")
	reader := bufio.NewReader(os.Stdin)
	longUrl, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	longUrl = strings.TrimSuffix(longUrl, "\n")

	client := resty.New()

	resp, err := client.R().
		SetFormData(map[string]string{
			"url": longUrl,
		}).
		Post(endpoint)

	if err != nil {
		panic(err)
	}

	fmt.Println("status-code:", resp.StatusCode())
	fmt.Println("response:")
	fmt.Println(resp.String())
}
