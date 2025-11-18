package resizeUrl

import (
	"shortener/test"
	"testing"
)

func TestResizeUrl(t *testing.T) {

	type TestWantResizeUrl struct {
		WantText string
		TestText string
	}

	tests := []test.TestCommon[TestWantResizeUrl]{
		{
			Name: "resizeUrl service url test #1",
			Want: TestWantResizeUrl{
				WantText: "05046f26",
				TestText: "https://google.com",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if got := ResizeUrl(tt.Want.TestText); got != tt.Want.WantText {
				t.Errorf("%s: ResizeUrl() = %s; want = %s", tt.Name, ResizeUrl(tt.Want.TestText), tt.Want.WantText)
			}
		})
	}
}
