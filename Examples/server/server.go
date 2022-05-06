package main
 
import (
   "fmt"
   "net/http"
   "httpserver/pkg"
)
 
func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
	fmt.Println("hello")
	fmt.Fprintf(w, pkg.WelcomeText)
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		fmt.Fprintf(w, "%v: %v\n", name, headers)
		fmt.Fprintf(w, "%v: %v\n", name, headers)
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8789", nil)
}