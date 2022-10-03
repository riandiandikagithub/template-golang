package core

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
)

var (
	readtimeout       string
	readheadertimeout string
	writetimeout      string
	idletimeout       string
	maxheaderbyte     string

	serverfileca         string
	serverfileprivatekey string
	serverfilepubkey     string
	servertls12client    string

	consuladdres             string
	consulregid              string
	consulregname            string
	consulregserver          string
	consulregport            string
	consulhealtCheckHttp     string
	consulhealtcheckInterval string
	consulhealtcheckTimeout  string
	consulonof               string

	kafkaBrokerUrl            string
	kafkaClient               string
	kafkaProducerTimeout      string
	kafkaProducerDialTimeout  string
	kafkaProducerReadTimeout  string
	kafkaProducerWriteTimeout string
	kafkaProducerMaxmsgbyte   string
)

const (
	CONSULREGPREFIX     = "consulreg"
	SERVERPREFIX        = "server"
	KAFKAPRODUCERPREFIX = "kafkapro"
	KAFKACONSUMERPREFIX = "kafkaconsumer"
	OTTOHTTPCLIENT      = "ottohttpclient"
)

type ServerConfig struct {

	// ReadTimeout is the maximum duration for reading the entire
	// request, including the body.
	//
	// Because ReadTimeout does not let Handlers make per-request
	// decisions on each request body's acceptable deadline or
	// upload rate, most users will prefer to use
	// ReadHeaderTimeout. It is valid to use them both.
	ReadTimeout time.Duration

	// ReadHeaderTimeout is the amount of time allowed to read
	// request headers. The connection's read deadline is reset
	// after reading the headers and the Handler can decide what
	// is considered too slow for the body.
	ReadHeaderTimeout time.Duration `envconfig:"rhtimeout"`

	// WriteTimeout is the maximum duration before timing out
	// writes of the response. It is reset whenever a new
	// request's header is read. Like ReadTimeout, it does not
	// let Handlers make decisions on a per-request basis.
	WriteTimeout time.Duration

	// IdleTimeout is the maximum amount of time to wait for the
	// next request when keep-alives are enabled. If IdleTimeout
	// is zero, the value of ReadTimeout is used. If both are
	// zero, ReadHeaderTimeout is used.
	IdleTimeout time.Duration

	// MaxHeaderBytes controls the maximum number of bytes the
	// server will read parsing the request header's keys and
	// values, including the request line. It does not limit the
	// size of the request body.
	// If zero, DefaultMaxHeaderBytes is used.
	MaxHeaderBytes int `envconfig:"maxbytes"`

	Serverfileca         string `envconfig:"FILECA"`
	Serverfileprivatekey string `envconfig:"PRIVATEKEY"`
	Serverfilepubkey     string `envconfig:"PUBLICKEY"`
	Servertls12client    string `envconfig:"TLS12STATUS"`
}

func GetServerTlsConfig() *tls.Config {
	if servertls12client == "ON" {

		caCert, err := ioutil.ReadFile(serverfileca)
		if err != nil {
			fmt.Println("Error : ", err)

		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		tlsConfig := &tls.Config{
			ClientCAs:  caCertPool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		}
		tlsConfig.BuildNameToCertificate()
		return tlsConfig
	}
	return &tls.Config{InsecureSkipVerify: true}

}

func GetServerConfig() *ServerConfig {
	var serverCfg ServerConfig
	err := envconfig.Process(SERVERPREFIX, &serverCfg)
	fmt.Println("Error Config Consul : ", err)
	return &serverCfg
}

func GinServerUp(listenAddr string, router *gin.Engine) error {

	cfg := *GetServerConfig()
	fmt.Println("[TLS.1.2]:", cfg.Servertls12client)
	srv := &http.Server{
		Addr:              listenAddr,
		Handler:           router,
		TLSConfig:         GetServerTlsConfig(),
		ReadTimeout:       cfg.ReadTimeout,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
		WriteTimeout:      cfg.WriteTimeout,
		IdleTimeout:       cfg.IdleTimeout,
		MaxHeaderBytes:    cfg.MaxHeaderBytes,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
	}

	if cfg.Servertls12client == "ON" {
		return srv.ListenAndServeTLS(cfg.Serverfilepubkey, cfg.Serverfileprivatekey)
	}
	return srv.ListenAndServe()
}
