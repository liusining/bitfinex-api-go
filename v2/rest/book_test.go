package rest

import (
	"bytes"
	"github.com/liusining/bitfinex-api-go/v2"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestBookAll(t *testing.T) {
	httpDo := func(_ *http.Client, req *http.Request) (*http.Response, error) {
		msg := `[[10579,1,0.0329596],[10578,1,0.11030234],[10577,2,0.11890895],[10576,2,1.0427],[10574,2,0.98962806],[10573,1,0.9443],[10572,1,0.06824617],[10571,1,0.42609023],[10570,1,0.002],[10569,2,0.99085269],[10568,3,2.1616],[10567,1,0.49990559],[10566,1,0.5],[10565,1,0.5413],[10564,2,0.99990599],[10563,2,0.28270321],[10561,2,0.99896343],[10560,1,0.498983],[10559,3,1.43741793],[10558,4,1.17],[10557,2,2.42],[10556,3,4.25833255],[10555,4,6.472],[10554,1,0.2],[10553,2,0.06940968],[10580,3,-4.5235],[10581,1,-0.9452],[10584,2,-0.46850263],[10585,1,-0.01],[10586,2,-0.93153],[10587,3,-0.82382839],[10589,2,-0.56565545],[10590,2,-0.43420271],[10592,1,-0.1],[10593,3,-2.1],[10594,3,-19.47635006],[10595,4,-7.352],[10596,1,-1.5],[10597,1,-4.5],[10598,1,-2.96],[10600,3,-0.41500001],[10601,1,-0.02835606],[10606,3,-0.28310301],[10607,2,-0.99729895],[10608,3,-0.25],[10609,3,-2.04831264],[10610,1,-0.05],[10613,2,-2],[10614,1,-0.3],[10615,1,-0.002]]`
		resp := http.Response{
			Body:       ioutil.NopCloser(bytes.NewBufferString(msg)),
			StatusCode: 200,
		}
		return &resp, nil
	}

	book, err := NewClientWithHttpDo(httpDo).Book.All("tBTCUSD", bitfinex.Precision0, 25)

	if err != nil {
		t.Fatal(err)
	}

	if len(book.Snapshot) != 50 {
		t.Fatalf("expected 50 book update entries in snapshot, but got %d", len(book.Snapshot))
	}
}
