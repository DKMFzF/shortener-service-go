package getUrl

import (
	"io"
	"net/http"
	"net/http/httptest"
	"shortener/test"
	"testing"

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
			r := httptest.NewRequest(http.MethodGet, tt.Want.URL, nil)

			// add param in url
			r.SetPathValue("id", tt.Want.URL[1:])
			r.Header.Set("Content-Type", "text/plain; charset=utf-8")

			w := httptest.NewRecorder()

			GetUrlHandler(w, r)
			res := w.Result()

			assert.Equal(t, tt.Want.Code, res.StatusCode)
			defer res.Body.Close()

			resBody, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("Error in read resBody in handler: %v", err)
				return
			}

			require.Equal(t, tt.Want.Response, string(resBody))
			assert.Equal(t, res.Header.Get("Content-Type"), tt.Want.ContentType)
		})
	}
}
