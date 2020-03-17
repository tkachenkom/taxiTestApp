package config

import (
	"github.com/tkachenkom/taxiTestApp/db"
	"sync"

	"github.com/sirupsen/logrus"
)

// IConfig general config interface
type IConfig interface {
	HTTP() *HTTP
	Log() *logrus.Entry
	DB() db.QInterface
}

// Config implementation of IConfig
type Config struct {
	sync.Mutex

	//internal objects
	http *HTTP
	log  *logrus.Entry
	db   db.QInterface
}

// New method returns a new IConfig for the app
func New() IConfig {
	return &Config{
		Mutex: sync.Mutex{},
	}
}
