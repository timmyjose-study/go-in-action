package listing01

import (
	"net/http"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestDownload(t *testing.T) {
	url := "https://timmyjose.github.io"
	statusCode := 200

	t.Log("Given to need to test for downloading content from a url")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"", url, statusCode)
		{
			resp, err := http.Get(url)
			if err != nil {
				t.Fatal("\t\tShould be able to make the Get call", ballotX, err)
			}
			t.Log("\t\tShould be able to make the Get call", checkMark)

			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t\tShould receive a \"%d\" status code. %v", statusCode, checkMark)
			} else {
				t.Errorf("\t\tShould  receive a \"%d\" statuc code. %v", statusCode, ballotX)
			}
		}
	}

}