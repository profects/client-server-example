module git.profects.com/profects/server

go 1.13

// Pin go-micro on v1.8.3
replace github.com/micro/go-micro => github.com/micro/go-micro v1.7.0

// Fix github.com/go-resty/resty unknown revision
replace github.com/go-resty/resty => gopkg.in/resty.v1 v1.11.0

// Fix github.com/gogo/protobuf@v0.0.0-20190410021324-65acae22fc9: invalid pseudo-version: revision is shorter than canonical (65acae22fc9d)
replace github.com/gogo/protobuf v0.0.0-20190410021324-65acae22fc9 => github.com/gogo/protobuf v1.2.2-0.20190723190241-65acae22fc9d

// Fix not enough arguments in call to watch.NewStreamWatcher
replace k8s.io/client-go => k8s.io/client-go v0.0.0-20190620085101-78d2af792bab

// Below replaces needed for go-micro v1.7.0
replace github.com/micro/go-plugins => github.com/micro/go-plugins v1.0.0

replace github.com/census-instrumentation/opencensus-proto v0.1.0-0.20181214143942-ba49f56771b8 => github.com/census-instrumentation/opencensus-proto v0.0.3-0.20181214143942-ba49f56771b8

replace github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1

// Fix module github.com/go-log/log@latest (v0.1.0) found, but does not contain package github.com/go-log/log/print
replace github.com/go-log/log => github.com/go-log/log v0.1.1-0.20181211034820-a514cf01a3eb

require (
	github.com/golang/protobuf v1.3.2
	github.com/micro/examples v0.2.0
	github.com/micro/go-plugins v1.1.2-0.20190710094942-bf407858372c
	golang.org/x/net v0.0.0-20190912160710-24e19bdeb0f2
	google.golang.org/grpc v1.23.1
)
