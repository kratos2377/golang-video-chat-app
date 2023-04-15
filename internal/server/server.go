package server

import (
	"flag"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"github.com/gofiber/websocket/v2"
	"golang.org/x/net/websocket"
)

var (
	addr = flag.String("addr,":"", os.Getenv("PORT"),"")
	cert = flag.String("cert", "", "")
	key = flag.String("key", "" , "")
)


func Run() error {
	flag.Parse()


	if *addr == ":" {
		*addr = ":8080"
	}


	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{Views: engine})
	app.User(logger.New())
	app.Use(Cors.New())
	app.Get("/" , handlers.Welcome)
	app.Get("/room/create", handlers.RoomCreate)
	app.Get("/room/:uuid", handlers.Room)
	app.Get("/room/:uuid/websocker", websocket.New(handlers.RoomWebSocket, websocket.Config{
		HandshakeTimeout: 10*time.Second
	}))
	app.Get("/room/:uuid/chat", handlers.RoomChat)
	app.Get("/room/:uuid/chat/websocket", websocket.New(handlers.RoomChatWebsocket))
	app.Get("/room/uuid/viewer/websocket", websocket.New(handlers.RoomViewerWebsocket))
	app.Get("/stream/:ssuid" , handlers.Stream)
	app.Get("/stream/:ssuid/websocket",)
	app.Get("/stram/ssuid/chat/websocket")

}