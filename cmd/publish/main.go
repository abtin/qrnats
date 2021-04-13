package main

import (
	"flag"
	"fmt"
	"github.com/abtin/qrnats/internal/model"
	"github.com/nats-io/nats.go"
	"os"
)

func main() {
	var (
		jsonFile string
		subject  string
	)
	usage := "publish -f <jsonfile> -s <subject>"
	flag.StringVar(&jsonFile, "f", "", usage)
	flag.StringVar(&subject, "s", "", usage)
	flag.Parse()
	if jsonFile == "" || subject == "" {
		fmt.Println(usage)
		os.Exit(1)
	}
	file, err := os.Open(jsonFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	user, err := model.NewUserFromJson(file)
	if err != nil {
		fmt.Println(err)
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
	if err := nc.Publish(subject, []byte(user.String())); err != nil {
		fmt.Printf("Error publishing message - %s\n", err)
	}
}
