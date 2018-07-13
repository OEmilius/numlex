package webserver

import (
	"fmt"
	"log"
	"net/http"
	"numlex"
)

var CDMN *numlex.CDMN

type Server struct {
	Addr string
}

func NewServer(addr string) (s Server) {
	s.Addr = addr
	return s
}

func (s *Server) StartWebServer() {
	http.HandleFunc("/", serveHome)
	fmt.Println("ListenAndServe:", s.Addr)
	err := http.ListenAndServe(s.Addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr, r.Method, r.URL)
	keys, ok := r.URL.Query()["num"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		w.Write([]byte("example usage http://localhost:8080/?num=9000000030"))
	} else {
		num := keys[0]
		log.Println("Url Param num is: " + string(num))
		answer := CDMN.GetRNStep(num)
		log.Println(answer)
		w.Write([]byte(answer))
	}
}
