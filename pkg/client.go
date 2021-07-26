package pkg

import (
	"github.com/raf924/bot/pkg/queue"
	"github.com/raf924/bot/pkg/users"
	"github.com/raf924/connector-api/pkg/gen"
	"google.golang.org/protobuf/proto"
)

type BotClientRelay struct {
	users    *users.UserList
	producer *queue.Producer
	consumer *queue.Consumer
	botUser  *gen.User
	commands []*gen.Command
}

func (b *BotClientRelay) GetUsers() []*gen.User {
	return b.users.All()
}

func (b *BotClientRelay) OnUserJoin(f func(user *gen.User, timestamp int64)) {
}

func (b *BotClientRelay) OnUserLeft(f func(user *gen.User, timestamp int64)) {
}

func (b *BotClientRelay) Connect(registration *gen.RegistrationPacket) (*gen.User, error) {
	b.commands = registration.Commands
	return b.botUser, nil
}

func (b *BotClientRelay) Send(packet *gen.BotPacket) error {
	return b.producer.Produce(packet)
}

func (b *BotClientRelay) Recv() (proto.Message, error) {
	m, err := b.consumer.Consume()
	return m.(proto.Message), err
}

func (b *BotClientRelay) Done() <-chan struct{} {
	return make(chan struct{})
}
