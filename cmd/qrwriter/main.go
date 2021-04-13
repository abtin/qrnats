package main

import (
	"flag"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/skip2/go-qrcode"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var (
		subject string
	)
	flag.StringVar(&subject, "s", "", "qrwriter -s <subject>")
	flag.Parse()
	if subject == "" {
		fmt.Println("qrwriter -s <subject>")
		os.Exit(1)
	}

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer ec.Close()

	_, err = ec.Subscribe(subject, listener)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("press ^C to break listening for messages")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGTERM, os.Interrupt)
	<-sc

	nc.Close()

}

func listener(m *nats.Msg) {
	now := time.Now()
	filename := fmt.Sprintf("qr-%d.png", now.Unix())
	err := qrcode.WriteFile(string(m.Data), qrcode.Medium, 256, filename)
	if err != nil {
		fmt.Println(err)
	}
}
