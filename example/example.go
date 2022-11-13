package example

import (
	"fmt"
	"net/http"
)

type home struct {
}

func (h *home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my homepage"))
}

func main() {
	h := home{}
	http.Handle("/", &h)

	fmt.Println("server running on 8080")
	http.ListenAndServe(":8080", nil)
}
