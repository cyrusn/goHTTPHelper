# golangHTTPHelper

golangHTTPHelper store some helper `func` for `http` package

- `Logger(handler http.Handler) http.Handler `
- `PrintJSON(w http.ResponseWriter, v interface{}) `
- `PrintError(w http.ResponseWriter, err error, errCode int) `
- `UnmarshalErrMessage(body []byte) (*ErrorMessage, error)`
