package p2p

import (
	"net/http"

	"github.com/ettoretoma/Nomad-coin-course/utils"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	_, err := upgrader.Upgrade(rw, r, nil)
	utils.HandleErr(err)
}
