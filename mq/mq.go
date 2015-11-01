package mq

import (
	"log"
	"os"
	"sync"
)

var consoleLog = log.New(os.Stdout, "[util.mq] ", log.LstdFlags)

type Provider interface {
	Name() string
	Dial(addr string, options ...string) error
	Subscribe(name string) (Subscription, error)
	Publish(name, message Message) error
	Close() error
}

type Subscription interface {
	Receive(receiveMsg func(messageStream chan Message, wg *sync.WaitGroup)) error
	Unsubscribe() error
}
