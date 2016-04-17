package util

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func SimpleJsonResponse(w http.ResponseWriter, statusCode int) {
	JsonResponse(w, nil, nil, statusCode)
}

func JsonResponse(w http.ResponseWriter, v interface{}, headers map[string]string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	//Any custom headers passed in
	for k, v := range headers {
		w.Header().Set(k, v)
	}
	w.WriteHeader(statusCode)
	if v != nil {
		b, _ := json.Marshal(v)
		fmt.Fprintf(w, "%s", string(b[:]))
	}
}


func ErrorResponse(w http.ResponseWriter, err error, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err != nil {
		b, _ := json.Marshal(map[string]interface{}{
			"message": err.Error(),
		})
		fmt.Fprintf(w, "%s", string(b[:]))
	}
}

func Jsonify(obj interface{}) string {
	b, err := json.MarshalIndent(obj, " ", "    ")
	if err != nil {
		return ""
	}
	return string(b)
}

func RemoveDuplicates(s []string) (t []string) {
	m := map[string]bool{}
	// walk the slice s and for each value we haven't seen so far, append it to t
	// this has the benefit of being clearer, not mutting the original underlyin array
	// as well as not hanging on to memory needlessly if the slice has few unique values
	for _, v := range s {
		if _, seen := m[v]; !seen {
			t = append(t, v)
			m[v] = true
		}
	}
	return t
}

func Random4DigitNumber() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	v := r.Intn(9999)
	if v == 0 {
		v = Random4DigitNumber()
	}
	return v
}
