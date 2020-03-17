package app

import (
	"crypto/rand"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/tkachenkom/taxiTestApp/config"
	"github.com/tkachenkom/taxiTestApp/db"
	"github.com/tkachenkom/taxiTestApp/db/models"
	"github.com/tkachenkom/taxiTestApp/server"
	"time"
)

// App parameters
type App struct {
	config config.IConfig
	log    *logrus.Entry
}

// New returns a new app with configured parameters
func New(config config.IConfig) *App {
	return &App{
		config: config,
		log:    config.Log(),
	}
}

// Start method will initialize and start the echo server
func (a *App) Start() error {
	conf := a.config
	httpConfiguration := conf.HTTP()

	e := echo.New()

	server.Router(
		conf.Log(),
		conf.DB(),
		e,
	)

	for i := 0; i < 50; i++ {
		err := create(conf.DB())
		if err != nil {
			return err
		}
	}

	serverHost := fmt.Sprintf("%s:%s", httpConfiguration.Host, httpConfiguration.Port)
	a.log.WithField("api", "start").
		Info(fmt.Sprintf("listenig addr =  %s, tls = %v", serverHost, httpConfiguration.SSL))

	if err := e.Start(serverHost); err != nil {
		return errors.Wrap(err, "failed to start http server")
	}

	go deleteOne(conf.DB())

	return nil
}

func generateName() string {
	b := make([]byte, 2)
	_, _ = rand.Read(b)

	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for i, m := range b {
		b[i] = letters[m%byte(len(letters))]
	}
	return string(b)
}

func create(db db.QInterface) error {
	id := uuid.New().String()

	err := db.OrdersQ().Insert(models.Order{
		ID:        id,
		Name:      generateName(),
		CreatedAt: time.Now(),
	})
	if err != nil {
		return errors.Wrap(err, "failed to insert new order to db")
	}

	return nil
}

func deleteOne(db db.QInterface) {
	for {
		time.Sleep(200 * time.Millisecond)

		err := db.OrdersQ().DeleteOne()
		if err != nil {
			logrus.Fatal(err)
		}

		err = create(db)
		if err != nil {
			logrus.Fatal(err)
		}
	}
}
