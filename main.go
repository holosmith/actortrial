package main

import (
	"fmt"
	"log"
	"time"

	"github.com/anthdm/hollywood/actor"
)

type Player struct {
	HP int
}

func NewPlayer(hp int) actor.Producer {
	return func() actor.Receiver {
		return &Player{
			HP: hp,
		}
	}
}

type TakeDamage struct {
	amount int
}

func (p *Player) Receive(c *actor.Context) {
	switch msg := c.Message().(type) {
	case actor.Started:
		fmt.Println("Reading the messages...")
	case actor.Stopped:
		fmt.Println("Tasks completed...")
	case TakeDamage:
		fmt.Println("Player is taking damage.", msg.amount)
	}
}

func main() {
	e, err := actor.NewEngine(actor.NewEngineConfig())
	if err != nil {
		log.Fatal(err)
	}

	pid := e.Spawn(NewPlayer(100), "player", actor.WithID("myuserid"))

	msg := TakeDamage{amount: 99}
	e.Send(pid, msg)

	time.Sleep(time.Second * 2)
}
