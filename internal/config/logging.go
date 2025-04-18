package config

import (
	"github.com/harluo/boot"
	"github.com/harluo/loki/internal/internal/kernel"
)

type Logging struct {
	// 日志级别
	Level string `default:"debug" json:"level,omitempty" valloggingate:"oneof=debug info warn error fatal"`
	// 类型
	Type kernel.Type `json:"type,omitempty" valloggingate:"required,oneof=loki"`
	// 日志调用方法过滤层级
	Skip int `default:"2" json:"skip,omitempty"`
	// 调用堆栈层级
	Stacktrace *int `json:"stacktrace,omitempty"`
	// Loki日志配置
	Loki *Loki `json:"loki,omitempty" valloggingate:"required_if=Type loki"`
}

func newLogging(config *boot.Config) (logging *Logging, err error) {
	logging = new(Logging)
	err = config.Build().Get(&struct {
		Logging *Logging `json:"logging,omitempty" valloggingate:"required"`
	}{
		Logging: logging,
	})

	return
}
