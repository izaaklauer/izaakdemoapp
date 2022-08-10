package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
)

const CONFIG_PATH_KEY = "CONFIG_PATH"

type config struct {
	HelloMessage string `yaml:"hello_message"`
	Port         int    `yaml:"port"`
}

var (
	HELLO_MESSAGE_ENV_KEY = "CONFIG_HELLO_MESSAGE"
	PORT_ENV_KEY          = "CONFIG_PORT"
)

func loadConfigFromFile() (*config, error) {

	configPath := os.Getenv(CONFIG_PATH_KEY)
	if configPath == "" {
		return nil, fmt.Errorf("%s not set", CONFIG_PATH_KEY)
	}
	configFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var c config
	if err := yaml.Unmarshal(configFile, &c); err != nil {
		return nil, fmt.Errorf("failed parsing config: %s", err)
	}

	return &c, nil
}

func loadConfigFromEnv() (*config, error) {
	helloMessage := os.Getenv(HELLO_MESSAGE_ENV_KEY)
	if helloMessage == "" {
		return nil, fmt.Errorf("Missing env var %q", HELLO_MESSAGE_ENV_KEY)
	}

	portStr := os.Getenv(PORT_ENV_KEY)
	if portStr == "" {
		return nil, fmt.Errorf("Missing env var %q", PORT_ENV_KEY)
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("Failed parsing port: %s", err)
	}

	return &config{
		HelloMessage: helloMessage,
		Port:         port,
	}, nil
}

func main() {
	cfg, err := loadConfigFromEnv()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("healthy\n"))
		return
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Demo Hello request from Addr: %s, Host: %s", r.RemoteAddr, r.Host)
		w.Write([]byte(fmt.Sprintf("Here is today's message (live demo): %q\n", cfg.HelloMessage)))
		return
	})

	// Apply default port if necessary
	var port int
	if cfg.Port != 0 {
		port = cfg.Port
	} else {
		port = 3001
		log.Printf("Defaulting to port %d", port)
	}

	serveAddr := fmt.Sprintf(":%d", port)
	log.Printf("Starting server at: %s\n", serveAddr)
	if err := http.ListenAndServe(serveAddr, nil); err != nil {
		log.Fatal(err)
	}
}
