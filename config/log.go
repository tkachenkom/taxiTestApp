package config

import (
	"github.com/caarlos0/env"
	"github.com/sirupsen/logrus"
)

// Log struct for config
type Log struct {
	Lvl string `env:"SIBEX_USER_PORTAL_LOG_LEVEL" envDefault:"debug"`
}

// GetLogEntry returns a new log entry
func (l *Log) GetLogEntry() *logrus.Entry {
	//err can be ignored in this case
	level, _ := logrus.ParseLevel(l.Lvl)

	logger := logrus.New()
	logger.SetLevel(level)

	return logrus.NewEntry(logger)
}

// Log entry
func (c *Config) Log() *logrus.Entry {
	if c.log != nil {
		return c.log
	}

	c.Lock()
	defer c.Unlock()

	log := &Log{}
	if err := env.Parse(log); err != nil {
		panic(err)
	}

	c.log = log.GetLogEntry()

	return c.log
}
