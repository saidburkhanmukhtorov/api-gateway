
prot-exp:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	export PATH="$PATH:$(go env GOPATH)/bin"

gen-proto-payment:
	protoc --go_out=genproto/ \
    --go-grpc_out=genproto/ \
	submodule/payment/*.proto


gen-proto-reservation:
	protoc --go_out=genproto/ \
    --go-grpc_out=genproto/ \
	submodule/reservation/*.proto
