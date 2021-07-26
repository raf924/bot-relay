package bot_relay

import (
	"github.com/raf924/bot-relay/pkg"
	"github.com/raf924/bot/pkg/relay/server"
)

func init() {
	server.RegisterRelayServer("bot", pkg.Builder)
}
