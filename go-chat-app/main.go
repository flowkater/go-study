package main

import (
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type conn struct {
	wsConn *websocket.Conn
	send   chan []byte
	wg     sync.WaitGroup
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		typ, reader, err := conn.NextReader()
		if err != nil {
			log.Println(err)
			return
		}

		writeCloser, err := conn.NextWriter(typ)
		if err != nil {
			log.Println(err)
			return
		}

		_, err = io.Copy(writeCloser, reader)
		if err != nil {
			log.Println(err)
			return
		}

		if err := writeCloser.Close(); err != nil {
			log.Println(err)
			return
		}
	}
}
func main() {
	if err := http.ListenAndServe(":8080", http.HandlerFunc(handler)); err != nil {
		log.Fatalln(err)
	}
}
