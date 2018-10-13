package errorlog

import (
	"log"
	"net/http"
)

// ErrorStatus ...
func ErrorStatus(w http.ResponseWriter, err error, status int) {
	if err != nil {
		log.Println("Error: ", err)
		w.WriteHeader(status)
		return
	}
}
