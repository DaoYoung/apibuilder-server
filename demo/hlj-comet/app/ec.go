package app

import (
	"github.com/nats-io/nats"
	"fmt"
	"hlj-comet/core/protocol"
)

var Ec *nats.EncodedConn

func InitEc() error {
	ds := fmt.Sprintf("nats://%s:%s@%s:%d",
		Config.Nats.User,
		Config.Nats.Password,
		Config.Nats.Host,
		Config.Nats.Port,
	)

	nc, err := nats.Connect(ds)
	if err != nil {
		return err
	}
	ec, err := nats.NewEncodedConn(nc, protocol.JsonIterator)
	if err != nil {
		return err
	}

	Ec = ec
	return nil
}
