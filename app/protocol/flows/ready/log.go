package ready

import (
	"github.com/Kash-Protocol/kashd/infrastructure/logger"
	"github.com/Kash-Protocol/kashd/util/panics"
)

var log = logger.RegisterSubSystem("PROT")
var spawn = panics.GoroutineWrapperFunc(log)
