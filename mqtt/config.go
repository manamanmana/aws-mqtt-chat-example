package mqtt

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"os"
)

type Config struct {
	Host       string `json:"host"`
	Port       int    `json:"port"`
	CaCert     string `json:"caCert"`
	ClientCert string `json:"clientCert"`
	PrivateKey string `json:"privateKey"`
}

func getSettingsFromFile(p string, opts *MQTT.ClientOptions) error {
	var (
		conf Config
		err  error
	)

	// Read condig json file
	conf, err = readFromConfigFile(p)
	if err != nil {
		log.SetOutput(os.Stderr)
		log.Error(err)
		return err
	}

	// Make TLS configulation
	var (
		tlsConfig *tls.Config
		ok        bool
	)
	tlsConfig, ok, err = makeTlsConfig(conf.CaCert, conf.ClientCert, conf.PrivateKey)
	if err != nil {
		return err
	}
	if ok {
		opts.SetTLSConfig(tlsConfig)
	}

	// Add Broker
	var brokerUri string = fmt.Sprintf("ssl://%s:%d", conf.Host, conf.Port)
	opts.AddBroker(brokerUri)

	return nil
}

func readFromConfigFile(path string) (Config, error) {
	var ret Config = Config{}

	var (
		b   []byte
		err error
	)

	b, err = ioutil.ReadFile(path)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(b, &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

func makeTlsConfig(cafile, cert, key string) (*tls.Config, bool, error) {
	var TLSConfig *tls.Config = &tls.Config{InsecureSkipVerify: false}
	var ok bool

	var certPool *x509.CertPool
	var err error
	var tlsCert tls.Certificate
	if cafile != "" {
		certPool, err = getCertPool(cafile)
		if err != nil {
			return nil, false, err
		}
		TLSConfig.RootCAs = certPool
		ok = true
	}
	if cert != "" {
		certPool, err = getCertPool(cert)
		if err != nil {
			return nil, false, err
		}
		TLSConfig.ClientAuth = tls.RequireAndVerifyClientCert
		TLSConfig.ClientCAs = certPool
		ok = true

	}
	if key != "" {
		if cert == "" {
			return nil, false, fmt.Errorf("key specified but cert is not specified")
		}
		tlsCert, err = tls.LoadX509KeyPair(cert, key)
		if err != nil {
			return nil, false, err
		}
		TLSConfig.Certificates = []tls.Certificate{tlsCert}
		ok = true
	}
	return TLSConfig, ok, nil
}

func getCertPool(pemPath string) (*x509.CertPool, error) {
	var certs *x509.CertPool = x509.NewCertPool()
	var pemData []byte
	var err error

	pemData, err = ioutil.ReadFile(pemPath)
	if err != nil {
		return nil, err
	}
	certs.AppendCertsFromPEM(pemData)
	return certs, nil
}
