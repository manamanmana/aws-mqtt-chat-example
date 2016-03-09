package main

import (
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/manamanmana/aws-mqtt-chat-example/mqtt"
	"os"
)

var args *mqtt.ArgOption

func init() {
	args = &mqtt.ArgOption{}

	flag.StringVar(&args.Topic, "topic", "", "Topic name to subscribe and publish")
	flag.IntVar(&args.Qos, "qos", 0, "QoS of the topic communication.")
	flag.StringVar(&args.Conf, "conf", "", "Config file JSON path and name for accessing to AWS IoT endpoint")
	flag.StringVar(&args.ClientId, "client-id", "", "client id to connect with")

	if args.Topic == "" {
		log.SetOutput(os.Stderr)
		msg := "Please specify topic with --topic option."
		log.Error(msg)
		fmt.Fprintln(os.Stderr, msg)
		os.Exit(1)
	}

	if args.Conf == "" {
		log.SetOutput(os.Stderr)
		msg := "Please specify Config file path with --conf option."
		log.Error(msg)
		fmt.Fprintln(os.Stderr, msg)
		os.Exit(1)
	}
}

func input(pub chan string) {
	for {
		var input string
		fmt.Scanln(&input)
		pub <- input
	}
}

func main() {
	opts, err := mqtt.NewOption(args)
	if err != nil {
		log.SetOutput(os.Stderr)
		log.Error(err)
		fmt.Fprintf(os.Stderr, "Error on making client options: %s", err)
		os.Exit(2)
	}

	client := mqtt.NewMQTTClient(opts)
	defer client.Client.Disconnect(250)

	_, err = client.Connect()
	if err != nil {
		log.SetOutput(os.Stderr)
		log.Error(err)
		fmt.Fprintf(os.Stderr, "Error on client coneection: %s", err)
		os.Exit(3)
	}

	go client.Subscribe(args.Topic, args.Qos)

	go input(client.PubChan)

	for {
		select {
		case s := <-client.SubChan:
			msg := string(s.Payload())
			fmt.Printf("\nmsg:%s\n", msg)
		case p := <-client.PubChan:
			client.Publish(args.Topic, args.Qos, p)
		}
	}

}
