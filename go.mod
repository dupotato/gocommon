module gocommon

go 1.14

require (
	github.com/ghodss/yaml v1.0.0
	github.com/hyperledger/fabric-sdk-go v1.0.0
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/rs/zerolog v1.23.0
	github.com/spf13/cobra v1.2.1
	github.com/spf13/viper v1.8.1
	google.golang.org/grpc v1.39.1
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/dupotato/gocommon v0.0.1 => ./
