package adapter

import "net/http"

// ErrorStatus ...
func ErrorStatus(w http.ResponseWriter, err error, status int) {
	if err != nil {
		w.WriteHeader(status)
		return
	}
}
