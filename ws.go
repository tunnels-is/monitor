package main

import (
	"log"
	"time"

	"github.com/lxzan/gws"
)


type WebSocketHandler struct{}


func (h *WebSocketHandler) OnOpen(conn *gws.Conn) {
	log.Println("Connection opened")
}

func (h *WebSocketHandler) OnMessage(conn *gws.Conn, message *gws.Message) {
	log.Println("Received message:", string(message.Data.String()))

	err := conn.WriteMessage(gws.OpcodeText, []byte("Hello, World!"))
	if err != nil {
		log.Println("Write error:", err)
	}
}

func (h *WebSocketHandler) OnClose(conn *gws.Conn, err error) {
	log.Println("Connection closed:", err)
}

func (c *WebSocketHandler) OnPing(socket *gws.Conn, payload []byte) {
	_ = socket.SetDeadline(time.Now().Add(PingInterval + PingWait))
	_ = socket.WritePong(nil)
}

func (c *WebSocketHandler) OnPong(socket *gws.Conn, payload []byte) {}


const (
	PingInterval = 5 * time.Second
	PingWait     = 10 * time.Second
)

type Handler struct{}

func (c *Handler) OnOpen(socket *gws.Conn) {
	_ = socket.SetDeadline(time.Now().Add(PingInterval + PingWait))
}

func (c *Handler) OnClose(socket *gws.Conn, err error) {}

func (c *Handler) OnPing(socket *gws.Conn, payload []byte) {
	_ = socket.SetDeadline(time.Now().Add(PingInterval + PingWait))
	_ = socket.WritePong(nil)
}

func (c *Handler) OnPong(socket *gws.Conn, payload []byte) {}

func (c *Handler) OnMessage(socket *gws.Conn, message *gws.Message) {
	defer message.Close()
	socket.WriteMessage(message.Opcode, message.Bytes())
}