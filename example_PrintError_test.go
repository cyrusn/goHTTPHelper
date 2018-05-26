package helper_test

import (
	"errors"
	"net/http"

	helper "github.com/cyrusn/goHTTPHelper"
)

func ExamplePrintError(w http.ResponseWriter, r *http.Request) {
	err := errors.New("invalid information")
	errCode := http.StatusBadRequest
	helper.PrintError(w, err, errCode)
}
