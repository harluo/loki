package core

import (
	"github.com/goexl/exception"
	"github.com/goexl/gox/field"
	"github.com/goexl/log"
	"github.com/goexl/loki"
	"github.com/harluo/http"
	"github.com/harluo/loki/internal/config"
	"github.com/harluo/loki/internal/internal/kernel"
)

type Logger = log.Logger

func newLogger(config *config.Logging, http *http.Client) (logger log.Logger, err error) {
	builder := log.New().Level(log.ParseLevel(config.Level))
	if nil != config.Stacktrace {
		builder.Stacktrace(*config.Stacktrace)
	}

	switch config.Type {
	case kernel.TypeLoki:
		self := config.Loki
		factory := loki.New().Url(self.Url).Http(http)
		if "" != self.Username || "" != self.Password {
			factory.Username(self.Username)
			factory.Password(self.Password)
		}
		if nil != self.Batch {
			factory = factory.Batch().Size(self.Batch.Size).Wait(self.Batch.Wait).Build()
		}
		if 0 != len(self.Labels) {
			factory.Labels(self.Labels)
		}
		if "" != self.Tenant {
			factory.Tenant(self.Tenant)
		}
		logger, err = builder.Factory(factory.Build()).Build()
	default:
		err = exception.New().Message("不支持的日志收集器").Field(field.New("type", config.Type)).Build()
	}

	return
}
