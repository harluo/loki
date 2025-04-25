package config

import (
	"github.com/harluo/config"
	"github.com/harluo/loki/internal/internal/kernel"
)

type Logging struct {
	// 日志级别
	Level string `default:"debug" json:"level,omitempty" validate:"oneof=debug info warn error fatal"`
	// 类型
	Type kernel.Type `json:"type,omitempty" validate:"required,oneof=loki"`
	// 日志调用方法过滤层级
	Skip int `default:"2" json:"skip,omitempty"`
	// 调用堆栈层级
	Stacktrace *int `json:"stacktrace,omitempty"`
	// Loki日志配置
	Loki *Loki `json:"loki,omitempty" validate:"required_if=Type loki"`
}

func newLogging(getter *config.Getter) (logging *Logging, err error) {
	logging = new(Logging)
	err = getter.Get(&struct {
		Logging *Logging `json:"logging,omitempty" validate:"required"`
	}{
		Logging: logging,
	})

	return
}
