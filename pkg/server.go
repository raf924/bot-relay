package pkg

import (
	"github.com/raf924/connector-api/pkg/gen"
	"google.golang.org/protobuf/proto"
)

type BotServerRelay struct {
}

func (b *BotServerRelay) Start(botUser *gen.User, users []*gen.User, trigger string) error {
	panic("implement me")
}

func (b *BotServerRelay) Commands() []*gen.Command {
	panic("implement me")
}

func (b *BotServerRelay) Send(message proto.Message) error {
	panic("implement me")
}

func (b *BotServerRelay) Recv() (*gen.BotPacket, error) {
	panic("implement me")
}
