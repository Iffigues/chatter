package main

import (
	"github.com/gorilla/websocket"
	"github.com/gorilla/mux"
	"fmt"
	"net/http"
	"encoding/json"
	"time"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 1024 * 1024
)

func init(){
	salle := Router.PathPrefix("/ws").Subrouter()
	salle.Path("/connect/{Key}").Methods("GET").HandlerFunc(serveWs)
}

var (
	stand = make(map[string]*froom)
	stande = make(map[string]*froom)
)

type client struct {
	ws   *websocket.Conn
	send chan []byte
	name interface{}
	room *froom
}

type hub struct {
	clients    map[*client]bool
	broadcast  chan []byte
	register   chan *client
	unregister chan *client
	content []byte
	live bool
}


type froom struct {
	private  bool
	password string
	user []userws
	room hub
}

type res struct{
     Texte string     
}

func  (tt *froom)run() {
	for {
		select {
		case c := <-tt.room.register:
			tt.room.clients[c] = true
			c.send <- []byte(tt.room.content)
			break

		case c := <-tt.room.unregister:
			_, ok := tt.room.clients[c]
			if ok {
				delete(tt.room.clients, c)
				close(c.send)
			}
			break

		case m := <-tt.room.broadcast:
			tt.room.content = m
			tt.broadcastMessage()
			break
		}
	}
}

func  (tt *froom)broadcastMessage() {
	for c := range tt.room.clients {
		select {
		case c.send <- []byte(tt.room.content):
			break
		default:
			close(c.send)
			delete(tt.room.clients, c)
		}
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  maxMessageSize,
	WriteBufferSize: maxMessageSize,
	EnableCompression: true,
}

func serveWs(w http.ResponseWriter,r *http.Request) {
	vars := mux.Vars(r)
	Key := vars["Key"]
	tt := stand[Key]
	session, _ := store.Get(r, "session-name")
	if !session.Values["co"].(bool){
		http.Error(w,"not connected",502)
		return
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if writeError(err){
		http.Error(w,"not a success",404)
		return
	}
	if tt.room.live == false{
		tt.room.live = true
	}
	c := &client{
		send: make(chan []byte, maxMessageSize),
		ws:   ws,
		name: session.Values["user"],
		room: tt,
	}
	tt.room.register <- c
	go c.writePump()
	c.readPump()
}

func (c *client) readPump() {
	defer func() {
		c.room.room.unregister <- c
		c.ws.Close()
	}()

	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error {
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println(message)
		c.room.room.broadcast <- message
	}
}

func (c *client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
 				c.write( websocket.CloseMessage,[]byte{},0)
				return
			}
			if err := c.write(websocket.TextMessage, message,1); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write( websocket.PingMessage,[]byte{},2); err != nil {
				return
			}
		}
	}
}

func (c *client) write(mt int, message []byte,yalta int) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	fr := res{}
	json.Unmarshal([]byte(message),&fr)
	//mimi,_ :=json.Marshal(jj(string(fr.Texte),c.name,yalta))
	//zz:=string(mimi)
	zz:="zzzz"
	return c.ws.WriteMessage(mt,[]byte(zz))
}
