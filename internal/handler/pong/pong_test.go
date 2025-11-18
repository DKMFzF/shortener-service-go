package pong

import (
	"shortener/test"

	"io"
	"net/http"
	"net/http/httptest"
	"testing"

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
			request := httptest.NewRequest(http.MethodGet, "/ping", nil)
			w := httptest.NewRecorder()
			PongHandler(w, request)

			res := w.Result()

			assert.Equal(t, tt.Want.Code, res.StatusCode)
			defer res.Body.Close()

			resBody, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("Error in io.read for body: %v", err)
				return
			}

			require.Equal(t, tt.Want.Response, string(resBody))
			assert.Equal(t, res.Header.Get("Content-Type"), tt.Want.ContentType)
		})
	}
}
