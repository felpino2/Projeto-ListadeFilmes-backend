package HTTP

import (
	DM "awesomeProject/psbackllfa/DataModel"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

func RunService() {
	s := &http.Server{
		Addr:         "localhost:8080",
		Handler:      Router{},
		ReadTimeout:  10 * time.Second, //request
		WriteTimeout: 10 * time.Second, //response
	}

	log.Fatal(s.ListenAndServe())
}

type Router struct{}

func (Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	var u DM.User
	body, _ := io.ReadAll(req.Body)
	json.Unmarshal(body, &u)
	//Insert(u, "users")
}
