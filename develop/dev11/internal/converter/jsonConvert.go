package converter

import (
	"encoding/json"
	"net/http"
)

type Result struct {
	Msg interface{} `json:"result"`
}

type Error struct {
	Err string `json:"error"`
}

func JsonResponse(w http.ResponseWriter, status int, msg interface{}, isError bool) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if isError {
		e := Error{
			msg.(string),
		}
		err := json.NewEncoder(w).Encode(e)
		if err != nil {
			http.Error(w, "Json encode error", http.StatusInternalServerError)
		}
	} else {
		res := Result{
			msg,
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			http.Error(w, "Json encode error", http.StatusInternalServerError)
		}
	}
}
