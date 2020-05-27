package config

import (
	"encoding/json"
	"io/ioutil"
	"log"

	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

// cfg is instance of config.
var (
	cfg Config
)

//FromFile parse config from config file
func FromFile(path string) (*Config, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}

	// app logging to file
	log.SetOutput(&lumberjack.Logger{
		Filename:   "app.log",
		MaxSize:    2, // megabytes
		MaxBackups: 5,
		MaxAge:     40,   //days
		Compress:   true, // disabled by default
	})

	log.Println("-------- * ------- Logging -------- * -------")
	return &cfg, nil
}

// Peek provides secure access to config options.
func Peek() *Config {
	return &cfg
}

// Config holds all config info.
type Config struct {
	Service  service  `json:"service"`  // Service holds service info
	Database database `json:"database"` // Database contains a dataaccess info
	Gatepay  gatepay  `json:"gatepay"`
	Nats     nats     `json:"nats"` // nats config params
}

type service struct {
	Name             string `json:"name"`
	Addr             string `json:"addr"`
	OutKey           string `json:"outkey"`
	InKey            string `json:"inkey"`
	Path             string `json:"path"`
	UploadedLogsPath string `json:"uploadedLogsPath"`
}

// Database holds dataaccess info.
type database struct {
	Addr   string `json:"addr"`
	User   string `json:"user"`
	Pass   string `json:"pass"`
	DbName string `json:"dbname"`
}

// gatepayment configs
type gatepay struct {
	Addr string `json:"addr"`
	User string `json:"user"`
	Pass string `json:"pass"`
}

type nats struct {
	Addr     string `json:"addr"`
	User     string `json:"user"`
	Pass     string `json:"pass"`
	CertFile string `json:"certFile"`
	CertKey  string `json:"certKey"`
	CertCa   string `json:"certCa"`
}
