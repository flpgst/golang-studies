package main

import "net/http"

func main() {
	http.HandleFunc("/", HelloWorldHandler)

	http.ListenAndServe(":8080", nil)

}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	param := r.URL.Query().Get("param")
	if param == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte("Hello World " + string(param)))
}
