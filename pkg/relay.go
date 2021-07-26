package pkg

import (
	"bytes"
	"github.com/raf924/bot/pkg"
	"github.com/raf924/bot/pkg/bot"
	bot2 "github.com/raf924/bot/pkg/config/bot"
	"github.com/raf924/bot/pkg/queue"
	client2 "github.com/raf924/bot/pkg/relay/client"
	"github.com/raf924/bot/pkg/relay/server"
	"github.com/raf924/bot/pkg/users"
	"github.com/raf924/connector-api/pkg/gen"
	"google.golang.org/protobuf/proto"
	"gopkg.in/yaml.v2"
)

type BotRelay struct {
	bot      pkg.Runnable
	botUser  *gen.User
	client   BotClientRelay
	users    *users.UserList
	producer *queue.Producer
	consumer *queue.Consumer
}

var Builder server.RelayServerBuilder = func(config interface{}) server.RelayServer {
	var buffer = bytes.NewBufferString("")
	if err := yaml.NewEncoder(buffer).Encode(config); err != nil {
		panic(err)
	}
	var cnf bot2.Config
	if err := yaml.NewDecoder(buffer).Decode(&cnf); err != nil {
		panic(err)
	}
	userList := users.NewUserList()
	client := BotClientRelay{
		users: userList,
	}
	client2.RegisterRelayClient("bot", func(config interface{}) client2.RelayClient {
		return &client
	})
	return &BotRelay{
		bot:    bot.NewBot(cnf),
		client: client,
		users:  userList,
	}
}

func (b *BotRelay) Start(botUser *gen.User, users []*gen.User, trigger string) error {
	b.client.botUser = botUser
	err := b.bot.Start()
	if err != nil {
		return err
	}
	for _, user := range users {
		b.users.Add(user)
	}
	return nil
}

func (b *BotRelay) Commands() []*gen.Command {
	return b.client.commands
}

func (b *BotRelay) Send(message proto.Message) error {
	return b.producer.Produce(message)
}

func (b *BotRelay) Recv() (*gen.BotPacket, error) {
	p, err := b.consumer.Consume()
	return p.(*gen.BotPacket), err
}
