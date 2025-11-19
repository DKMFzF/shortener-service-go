package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	endpoint := "http://localhost:8080/"
	data := url.Values{}

	fmt.Println("Введите длинный URL")
	reader := bufio.NewReader(os.Stdin)
	long, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	long = strings.TrimSuffix(long, "\n")
	data.Set("url", long)

	client := http.Client{}

	req, err := http.NewRequest(http.MethodPost, endpoint, strings.NewReader(data.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/x-www-from-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	fmt.Println("Статус-код ", res.Status)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
