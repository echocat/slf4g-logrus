package main

import (
	"errors"
	_ "github.com/echocat/slf4g-logrus/bridge/hook"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.WithField("a", 11).
		WithField("b", 12).
		Info("hello, world")

	logrus.WithField("a", 21).
		WithField("b", 22).
		WithError(errors.New("someError")).
		Error("hello,\nworld")
}
