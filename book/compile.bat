cd pkg/sdk
protoc -I . --go_out=.  --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative bookpb/service.proto
protoc -I . --go_out=.  --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative bookpb/book.proto
protoc -I . --go_out=.  --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative base/request_response.proto
cd ../..

xcopy pkg\sdk\bookpb\*.go ..\rpc\clients\rpc_book\bookpb\ /s /y