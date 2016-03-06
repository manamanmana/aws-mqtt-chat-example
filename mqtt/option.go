package mqtt

import (
	"crypto/rand"
	MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
	log "github.com/Sirupsen/logrus"
	"os"
)

var MaxClientIdLen = 8

type ArgOption struct {
	Topic    string
	Qos      int
	Conf     string
	ClientId string
	Host     string
	Port     int
	Cacert   string
	Cert     string
	Key      string
}

func NewOption(args *ArgOption) (*MQTT.ClientOptions, error) {
	var opts *MQTT.ClientOptions = MQTT.NewClientOptions()

	var host string = args.Host

	if host == "" {
		err := getSettingsFromFile(args.Conf, opts)
		if err != nil {
			log.SetOutput(os.Stderr)
			log.Error(err)
			return nil, err
		}
	}

	var clientId string = args.ClientId
	if clientId == "" {
		clientId = getRandomClientId()
	}
	opts.SetClientID(clientId)
	opts.SetAutoReconnect(true)
	return opts, nil
}

func getRandomClientId() string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, MaxClientIdLen)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return "mqttclient-" + string(bytes)
}
