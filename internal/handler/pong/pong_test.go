package pong

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPongHandler(t *testing.T) {
	type want struct {
		code        int
		response    string
		contentType string
	}

	tests := []struct {
		name string
		want want
	}{
		{
			name: "ping/pong test #1",
			want: want{
				code:        200,
				response:    "pong",
				contentType: "text/plain; charset=utf-8",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "/ping", nil)
			w := httptest.NewRecorder()
			PongHandler(w, request)

			res := w.Result()

			assert.Equal(t, tt.want.code, res.StatusCode)
			defer res.Body.Close()

			resBody, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("Error in io.read for body: %v", err)
				return
			}

			require.Equal(t, tt.want.response, string(resBody))
			assert.Equal(t, res.Header.Get("Content-Type"), tt.want.contentType)
		})
	}
}
