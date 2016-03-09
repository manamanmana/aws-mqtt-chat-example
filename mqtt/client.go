package mqtt

import (
	MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
	log "github.com/Sirupsen/logrus"
	"os"
)

type MQTTClient struct {
	Client  *MQTT.Client
	Opts    *MQTT.ClientOptions
	SubChan chan MQTT.Message
	PubChan chan string
}

func NewMQTTClient(opts *MQTT.ClientOptions) *MQTTClient {
	var client *MQTTClient = &MQTTClient{Opts: opts}
	client.SubChan = make(chan MQTT.Message)
	client.PubChan = make(chan string)

	return client
}

func (m *MQTTClient) Connect() (*MQTT.Client, error) {
	log.SetOutput(os.Stdout)
	m.Client = MQTT.NewClient(m.Opts)

	log.Info("Connecting to broler server.....")

	var token MQTT.Token = m.Client.Connect()
	if token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	return m.Client, nil
}

func (m *MQTTClient) Subscribe(topic string, qos int) error {
	log.SetOutput(os.Stdout)
	log.Infof("Start Subscribeing : %s ...", topic)

	var token MQTT.Token = m.Client.Subscribe(topic, byte(qos), func(client *MQTT.Client, msg MQTT.Message) {
		m.SubChan <- msg
	})

	if token.Wait() && token.Error() != nil {
		log.SetOutput(os.Stderr)
		log.Error(token.Error())
		return token.Error()
	}
	return nil
}

func (m *MQTTClient) Publish(topic string, qos int, input string) error {
	var token MQTT.Token = m.Client.Publish(topic, byte(qos), false, input)
	if token.Wait() && token.Error() != nil {
		log.SetOutput(os.Stderr)
		log.Error(token.Error())
		return token.Error()
	}
	return nil
}
