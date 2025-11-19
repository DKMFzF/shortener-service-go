package pong

import (
	"shortener/test"

	"net/http"
	"net/http/httptest"
	"testing"

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

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(PongHandler))
			defer ts.Close()

			client := resty.New()
			client.SetBaseURL(ts.URL)

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
