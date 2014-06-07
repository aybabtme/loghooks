package loghooks_test

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/aybabtme/loghooks"
	"io/ioutil"
	"time"
)

func ExampleHooks() {
	myHook := loghooks.Hook(func(e *logrus.Entry) error {
		tellTheWorld(e)
		return nil
	}, logrus.Info, logrus.Warn, logrus.Error)

	log := logrus.New()
	log.Out = ioutil.Discard
	log.Hooks.Add(myHook)

	log.Info("hello")
	log.Error("nothing here")
	log.Warn("oups")

	// Output:
	// hello
	// nothing here
	// oups
}

func tellTheWorld(e *logrus.Entry) {
	e.Data["time"] = time.Date(2000, 01, 01, 01, 01, 01, 01, time.UTC).String()
	fmt.Printf("%v\n", e.Data["msg"])
}
