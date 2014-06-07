/*
Package loghooks makes it easy to create custom hooks for logrus.

    myHook := loghooks.Hook(func(e *logrus.Entry) error {
        tellTheWorld(e)
        return nil
    }, logrus.Info, logrus.Warn, logrus.Error)

    log := logrus.New()
    log.Hooks.Add(myHook)

This saves you from creating whole new types when you just want a simple
little hook.
*/
package loghooks

import (
	"github.com/Sirupsen/logrus"
)

// Hook creates a hook type with your fire func and your levels. You don't
// need to implement a new type and you can create any hook you feel like,
// right there in your code.
func Hook(fire func(*logrus.Entry) error, levels ...logrus.Level) logrus.Hook {
	return hook{onFire: fire, levels: levels}
}

type hook struct {
	onFire func(*logrus.Entry) error
	levels []logrus.Level
}

func (h hook) Fire(entry *logrus.Entry) error { return h.onFire(entry) }
func (h hook) Levels() []logrus.Level         { return h.levels }
