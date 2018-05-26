package helper_test

import (
	"net/http"

	helper "github.com/cyrusn/goHTTPHelper"
)

func ExamplePrintJSON(w http.ResponseWriter, r *http.Request) {
	type Message struct {
		Body string
	}
	msg := Message{"Hello! World."}
	helper.PrintJSON(w, msg)
}
