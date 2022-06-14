module gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core

go 1.18

require gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons v0.0.0

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/go-ozzo/ozzo-validation/v4 v4.3.0
	gorm.io/datatypes v1.0.6
	gorm.io/gorm v1.23.4
)

replace gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons v0.0.0 => ../commons
