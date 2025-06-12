module example.com/xxx/interface

go 1.24.2

replace example.com/xxx/interface/api => ./api


require (
	example.com/xxx/interface/api v0.0.0
	example.com/xxx/user-service/api v0.0.0
	github.com/casbin/casbin/v2 v2.103.0
	github.com/deckarep/golang-set/v2 v2.8.0
	github.com/go-kratos/kratos/contrib/log/zap/v2 v2.0.0-20250314165958-d9aa7ff19541
	github.com/go-kratos/kratos/v2 v2.8.4
	github.com/google/wire v0.6.0
	github.com/gorilla/handlers v1.5.2
	go.uber.org/automaxprocs v1.6.0
	go.uber.org/zap v1.27.0
	go.uber.org/zap/exp v0.3.0
	google.golang.org/protobuf v1.36.6
	gorm.io/gorm v1.25.12
)

require (
	dario.cat/mergo v1.0.1 // indirect
	github.com/armon/go-metrics v0.4.1 // indirect
	github.com/bmatcuk/doublestar/v4 v4.8.1 // indirect
	github.com/casbin/govaluate v1.3.0 // indirect
	github.com/fatih/color v1.16.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.8.0 // indirect
	github.com/go-kratos/aegis v0.2.0 // indirect
	github.com/go-kratos/kratos/contrib/registry/consul/v2 v2.0.0-20250527152916-d6f5f00cf562 // indirect
	github.com/go-playground/form/v4 v4.2.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/hashicorp/consul/api v1.32.1 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.5.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/serf v0.10.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/exp v0.0.0-20250305212735-054e65f0b394 // indirect
	golang.org/x/net v0.38.0 // indirect
	golang.org/x/sync v0.12.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250603155806-513f23925822 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250528174236-200df99c418a // indirect
	google.golang.org/grpc v1.73.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
