package getUrl

import (
	"net/http/httptest"
	"shortener/test"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUrlHandler(t *testing.T) {
	tests := []test.TestCommon[test.TestWantWanthandlerWithUrl[string]]{
		{
			Name: "getUrl handler test #1",
			Want: test.TestWantWanthandlerWithUrl[string]{
				TestWantHandler: test.TestWantHandler[string]{
					Code:        200,
					ContentType: "text/plain; charset=utf-8",

					// TODO: write service logic for get url
					Response: "https://practicum.yandex.ru/",
				},
				URL: "/abs123",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)

			router := gin.New()
			router.GET("/:id", GetUrlHandler)

			ts := httptest.NewServer(router)
			defer ts.Close()

			client := resty.New()
			client.SetBaseURL(ts.URL)

			resp, err := client.R().
				SetHeader("Content-Type", "text/plain; charset=utf-8").
				Get(tt.Want.URL)

			if err != nil {
				panic(err)
			}

			require.Equal(t, tt.Want.Response, string(resp.Body()))
			assert.Equal(t, resp.Header().Get("Content-Type"), tt.Want.ContentType)
		})
	}
}
