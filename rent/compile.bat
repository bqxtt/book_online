cd pkg/sdk
protoc -I . --go_out=.  --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative userpb/service.proto
protoc -I . --go_out=.  --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative userpb/user.proto
protoc -I . --go_out=.  --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative base/request_response.proto
cd ../..

xcopy pkg\sdk\userpb\*.go ..\rpc\clients\rpc_user\userpb\ /s /y