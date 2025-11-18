package short

import (
	"io"
	"net/http"
	"net/http/httptest"
	"shortener/test"
	"strings"
	"testing"

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
			r := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.Want.URL))
			r.Header.Set("Content-Type", "text/plain; charset=utf-8")

			w := httptest.NewRecorder()

			ShortHandler(w, r)

			res := w.Result()

			assert.Equal(t, tt.Want.Code, res.StatusCode)
			defer res.Body.Close()

			resBody, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("Error read resBody: %v", err)
				return
			}

			require.Equal(t, tt.Want.Response, string(resBody))
			assert.Equal(t, tt.Want.ContentType, res.Header.Get("Content-Type"))
		})
	}
}
