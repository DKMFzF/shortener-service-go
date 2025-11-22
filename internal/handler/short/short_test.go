package short

import (
	"net/http/httptest"
	"shortener/test"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	baseUrl = "http://localhost:8080/"
)

func TestShortHandler(t *testing.T) {
	tests := []test.TestCommon[test.TestWantWanthandlerWithUrl[string]]{
		{
			Name: "short handler test #1",
			Want: test.TestWantWanthandlerWithUrl[string]{
				TestWantHandler: test.TestWantHandler[string]{
					Code:        201,
					ContentType: "text/plain; charset=utf-8",
					Response:    baseUrl + "05046f26",
				},
				URL: "https://google.com",
			},
		},
		{
			Name: "short handler test #2",
			Want: test.TestWantWanthandlerWithUrl[string]{
				TestWantHandler: test.TestWantHandler[string]{
					Code:        201,
					ContentType: "text/plain; charset=utf-8",
					Response:    baseUrl + "405c2462",
				},
				URL: "https://dkmfzf.com",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			router := gin.New()
			router.POST("/", ShortHandler)

			ts := httptest.NewServer(router)
			defer ts.Close()

			client := resty.New()
			client.SetBaseURL(ts.URL)

			resp, err := client.R().
				SetHeader("Content-Type", "text/plain; charset=utf-8").
				SetBody(tt.Want.URL).
				Post("/")

			if err != nil {
				panic(err)
			}

			require.Equal(t, tt.Want.Response, string(resp.Body()))
			assert.Equal(t, tt.Want.ContentType, resp.Header().Get("Content-Type"))
		})
	}
}
