package helper_test

import (
	"net/http"

	helper "github.com/cyrusn/goHTTPHelper"
	"github.com/gorilla/mux"
)

func ExampleLogger() {
	r := mux.NewRouter()
	address := "localhost:5000"
	http.ListenAndServe(address, helper.Logger(r))
}
