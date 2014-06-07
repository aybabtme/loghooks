# Package loghooks

A hook maker for [`logrus`](https://github.com/Sirupsen/logrus).

# Usage

```go
myHook := loghooks.Hook(func(e *logrus.Entry) error {
    tellTheWorld(e)
    return nil
}, logrus.Info, logrus.Warn, logrus.Error)

log := logrus.New()
log.Hooks.Add(myHook)
```
