package sample_test

import (
	"net/http"
	"testing"
)

func TestDownload(t *testing.T) {
	tt := []struct {
		url        string
		statusCode int
	}{
		{
			url:        "https://www.ardanlabs.com/blog/index.xml",
			statusCode: http.StatusOK,
		},
		{
			url:        "http://rss.cnn.com/rss/cnn_topstorie.rss",
			statusCode: http.StatusNotFound,
		},
	}

	for _, test := range tt {
		resp, err := http.Get(test.url)
		if err != nil {
			t.Fatalf("unable to issue GET on URL: %s: %s", test.url, err)
		}
		resp.Body.Close()

		if resp.StatusCode != test.statusCode {
			t.Log("exp:", test.statusCode)
			t.Log("got:", resp.StatusCode)
			t.Fatal("status code did not match up!")
		}
	}
}
