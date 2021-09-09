package main

import (
	"C"
)
import (
	"fmt"
	"math"
	"net/http"
)

//export pi
func pi() float64 {
	return math.Pi
}

//export serve
func serve(port int) {
	http.HandleFunc("/pi", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte(fmt.Sprintf("%f", pi())))
	})
	fmt.Printf("Starting server on port %d …\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func main() {}
