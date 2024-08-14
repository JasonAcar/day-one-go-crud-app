package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func ParseThisBody(r *http.Request, x interface{}) {
	d := json.NewDecoder(r.Body)
	err := d.Decode(&x)
	if err != nil {
		panic(err)
	}
	//err := json.Unmarshal([]byte(r.Body), &x)
}

func RespondWithError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	jsonResp := map[string]string{"error": message}
	json.NewEncoder(w).Encode(jsonResp)
}
