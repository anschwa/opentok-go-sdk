package opentok_test

import (
	"net/http"
	"time"

	"github.com/anschwa/opentok-go-sdk/v2/opentok"
)

const (
	apiKey    = "<your api key here>"
	apiSecret = "<your api secret here>"
)

var (
	ot     = opentok.New(apiKey, apiSecret, http.DefaultClient)
	client = &http.Client{Timeout: 120 * time.Second}
)

func ExampleNew() {
	const (
		apiKey    = "12345678"
		apiSecret = "ba7816bf8f01cfea414140de5dae2223b00361a3"
	)

	opentok.New(apiKey, apiSecret, client)
}
