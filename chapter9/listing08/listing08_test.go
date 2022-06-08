package listing08

import (
	"net/http"
	"testing"
)

const (
	checkMark = "\u2713"
	ballotX   = "\u2717"
)

func TestDownload(t *testing.T) {
	var urls = []struct {
		url        string
		statusCode int
	}{
		{
			"https://timmyjose.github.io",
			200,
		},
		{
			"https://z0ltan.wordpress.com",
			200,
		},
	}

	t.Log("Given the need to test download functionality for multiple urls")
	{
		for _, url := range urls {
			t.Logf("\tWhen checking \"%s\" for status code \"%d\"", url.url, url.statusCode)
			{
				resp, err := http.Get(url.url)
				if err != nil {
					t.Fatal("\t\tShould be able to Get the url", ballotX, err)
				}

				defer resp.Body.Close()

				if resp.StatusCode == url.statusCode {
					t.Logf("\t\tShould have a \"%d\" status code. %v", url.statusCode, checkMark)
				} else {
					t.Errorf("\t\tShould have a \"%d\" status code. %v", url.statusCode, ballotX)
				}
			}
		}
	}

}