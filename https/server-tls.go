package https

import (
	"fmt"
	"log"
	"net/http"
)

func Https() {
	http.HandleFunc("/", handler)
	log.Print("Start to listening the incoming requests on https address: 0.0.0.0:443\n")
	if err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil); err != nil {
		log.Print(err.Error())
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi")
}
