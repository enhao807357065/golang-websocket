package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"github.com/drone/routes"
	//	"log"
	"net/http"
)

func main() {
	fmt.Println("Start app")
	//	//websocket
	//	http.Handle("/", websocket.Handler(Echo))
	//	if err := http.ListenAndServe(":4001", nil); err != nil {
	//		log.Fatal("ListenAndServe:", err)
	//	}

	//restful
	mux := routes.New()
	mux.Get("/user/:uid", GetUser)
	http.Handle("/", mux)
	http.ListenAndServe(":4002", nil)
}

func Echo(ws *websocket.Conn) {
	var err error
	for {
		var reply string
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}
		fmt.Println("Received back from client: " + reply)
		msg := "Received:  " + reply
		fmt.Println("Sending to client: " + msg + "_server")
		if err = websocket.Message.Send(ws, msg+"_server"); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Println("the request uid is: ", uid)
}
