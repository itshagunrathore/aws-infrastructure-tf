
package services

import (
	"gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/log"
)

func HandleService() {
	logger := log.Logger()
	logger.Info("Hello World from service")
	logger.Error("Not able to reach blog. from service")
}
