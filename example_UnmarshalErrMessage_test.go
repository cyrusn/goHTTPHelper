package helper_test

import (
	"fmt"
	"io/ioutil"
	"net/http"

	helper "github.com/cyrusn/goHTTPHelper"
)

func ExampleUnmarshalErrMessage(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	errCode := http.StatusBadRequest
	if err != nil {
		helper.PrintError(w, err, errCode)
	}

	errMsg, err := helper.UnmarshalErrMessage(body)
	if err != nil {
		helper.PrintError(w, err, errCode)
	}

	fmt.Println(errMsg.Code, errMsg.Message)
}
