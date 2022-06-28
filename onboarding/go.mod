module gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/onboarding

go 1.18

require gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons v0.0.0

require gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons/dsa v0.0.0

require gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/core/src/repositories v0.0.0

require github.com/gin-gonic/gin v1.8.0

require gorm.io/datatypes v1.0.6

require gorm.io/gorm v1.23.4

require github.com/jmespath/go-jmespath v0.4.0 // indirect

require (
	github.com/aws/aws-sdk-go v1.44.34
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-ozzo/ozzo-validation/v4 v4.3.0 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.10.0 // indirect
	github.com/goccy/go-json v0.9.7 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.1 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.21.0 // indirect
	golang.org/x/crypto v0.0.0-20220525230936-793ad666bf5e // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace gitlab.teracloud.ninja/teracloud/pod-services/baas-spike/commons v0.0.0 => ../commons
