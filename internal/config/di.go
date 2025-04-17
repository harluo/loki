package config

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Get().Dependency().Puts(
		newLogging, // 配置
	).Build().Apply()
}
