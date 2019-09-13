module git.profects.com/profects/server

go 1.13

replace github.com/micro/go-plugins => github.com/micro/go-plugins v1.2.0

// Fix github.com/gogo/protobuf@v0.0.0-20190410021324-65acae22fc9: invalid pseudo-version: revision is shorter than canonical (65acae22fc9d)
replace github.com/gogo/protobuf v0.0.0-20190410021324-65acae22fc9 => github.com/gogo/protobuf v1.2.2-0.20190723190241-65acae22fc9d

require (
	github.com/golang/protobuf v1.3.2
	github.com/micro/examples v0.2.0
	github.com/micro/go-micro v1.10.0
	github.com/micro/go-plugins v1.1.2-0.20190710094942-bf407858372c
	golang.org/x/net v0.0.0-20190912160710-24e19bdeb0f2
	google.golang.org/grpc v1.23.1
)
