// Package helper stores some helper func for http package
package helper

// httpLogger is a simple logger to log the request info
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// ErrorMessage store error information
type ErrorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Logger is a middleware to log all http requests
func Logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

// PrintJSON print json string to http.ResponseWriter
func PrintJSON(w http.ResponseWriter, v interface{}) {
	errCode := http.StatusBadRequest
	if err := json.NewEncoder(w).Encode(v); err != nil {
		PrintError(w, err, errCode)
	}
}

// PrintError print error message as json string to http.ResponseWriter
func PrintError(w http.ResponseWriter, err error, errCode int) {
	errMessage := generateErrorMessage(errCode, err)
	http.Error(w, errMessage, errCode)
}

// UnmarshalErrMessage unmarshal error message
func UnmarshalErrMessage(body []byte) (*ErrorMessage, error) {
	errMessage := new(ErrorMessage)
	err := json.Unmarshal(body, errMessage)
	if err != nil {
		return nil, err
	}
	return errMessage, nil
}

func generateErrorMessage(code int, message error) string {
	b, _ := json.Marshal(ErrorMessage{code, message.Error()})
	return fmt.Sprintf("%s", b)
}
