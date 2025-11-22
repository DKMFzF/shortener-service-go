package pong

import (
	"shortener/test"

	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPongHandler(t *testing.T) {
	tests := []test.TestCommon[test.TestWantHandler[string]]{
		{
			Name: "ping/pong test #1",
			Want: test.TestWantHandler[string]{
				Code:        200,
				Response:    "pong",
				ContentType: "text/plain; charset=utf-8",
			},
		},
	}

	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.GET("/ping", PongHandler)

	ts := httptest.NewServer(router)
	defer ts.Close()

	client := resty.New()
	client.SetBaseURL(ts.URL)

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			resp, err := client.R().Get("/ping")
			if err != nil {
				panic(err)
			}

			assert.Equal(t, tt.Want.Code, resp.StatusCode())
			require.Equal(t, tt.Want.Response, string(resp.Body()))
			assert.Equal(t, tt.Want.ContentType, resp.Header().Get("Content-Type"))
		})
	}
}
