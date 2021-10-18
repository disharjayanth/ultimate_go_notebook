package sample_test

import (
	"net/http"
	"testing"
)

func TestDownload(t *testing.T) {
	tt := []struct {
		name       string
		url        string
		statusCode int
	}{
		{
			"ok",
			"https://www.ardanlabs.com/blog/index.xml",
			http.StatusOK,
		},
		{
			"notfound",
			"http://rss.cnn.com/rss/cnn_topstorie.rss",
			http.StatusNotFound,
		},
	}

	for _, test := range tt {
		test := test
		tf := func(t *testing.T) {
			// runs test in parllel
			// t.Parallel()
			resp, err := http.Get(test.url)
			if err != nil {
				t.Fatalf("unable to issue GET/URL: %s: %s", test.url, err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != test.statusCode {
				t.Log("exp:", test.statusCode)
				t.Log("got:", resp.StatusCode)
				t.Fatal("status code dont match!")
			}
		}

		t.Run(test.name, tf)
	}
}
