package socket

import (
	socketio "github.com/googollee/go-socket.io"
	"log"
)

var Manager *socketio.Server

func InitSocket(userID int64) *socketio.Server {
	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext(userID)
		log.Println("✅ User connected:", userID)
		return nil
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("❌ User disconnected:", s.Context())
	})

	Manager = server
	go server.Serve()
	return server
}

func EmitToUser(userID string, event string, data interface{}) {
	if Manager == nil {
		log.Println("⚠️ Socket manager not initialized")
		return
	}

	for _, conn := range Manager.GetNamespace("/") {
		if conn.Context() == userID {
			conn.Emit(event, data)
		}
	}
}