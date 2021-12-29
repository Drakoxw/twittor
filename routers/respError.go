package routers

import (
	"encoding/json"
	"net/http"
)

/*ResponseErr => responde con el Error amigable */
func ResponseErr(rw http.ResponseWriter, msgErr string, statusCode int) {
	data := HttpErr2{Err: true, Msg: msgErr}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(statusCode)
	json.NewEncoder(rw).Encode(data)
}
