package grpcclient

import (
	"github.com/Kash-Protocol/kashd/infrastructure/logger"
	"github.com/Kash-Protocol/kashd/util/panics"
)

var log = logger.RegisterSubSystem("RPCC")
var spawn = panics.GoroutineWrapperFunc(log)
