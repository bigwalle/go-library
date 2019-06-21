module go-library

go 1.12

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/dgryski/go-farm v0.0.0-20190323231341-8198c7b169ec
	github.com/go-ego/cedar v0.0.0-20190318172318-7949cd3f793f
	github.com/go-sql-driver/mysql v1.4.1
	github.com/gogo/protobuf v1.2.0
	github.com/golang/protobuf v1.2.0
	github.com/json-iterator/go v1.1.6
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/montanaflynn/stats v0.5.0
	github.com/naoina/go-stringutil v0.1.0 // indirect
	github.com/naoina/toml v0.1.1
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v0.9.2
	github.com/samuel/go-zookeeper v0.0.0-20180130194729-c4fab1ac1bec
	github.com/sirupsen/logrus v1.4.1
	github.com/smartystreets/goconvey v0.0.0-20190330032615-68dc04aab96a
	github.com/stretchr/testify v1.3.0
	github.com/tsuna/gohbase v0.0.0-20190201102810-d3184c1526df
	github.com/welcome112s/go-library v0.0.0-20190621050001-4a09ffe310de
	golang.org/x/net v0.0.0-20190311183353-d8887717615a
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4
	google.golang.org/grpc v1.20.1
	gopkg.in/go-playground/validator.v9 v9.26.0
)

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.26.0
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190123085648-057139ce5d2b
	golang.org/x/lint => github.com/golang/lint v0.0.0-20181026193005-c67002cb31c3
	golang.org/x/net => github.com/golang/net v0.0.0-20190311183353-d8887717615a
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20180821212333-d2e6202438be
	golang.org/x/sync => github.com/golang/sync v0.0.0-20181108010431-42b317875d0f
	golang.org/x/sys => github.com/golang/sys v0.0.0-20180905080454-ebe1bf3edb33
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190328211700-ab21143f2384
	google.golang.org/appengine => github.com/golang/appengine v1.1.0
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20180817151627-c66870c02cf8
	google.golang.org/grpc => github.com/grpc/grpc-go v1.20.1
)
