# MeroDB
Simple DB implementation

### Tasks
- [x] In Memory read write
- [ ] Compiler
- [x] FS Read/Write

## Generate protocol buffers
```bash
export PATH="$PATH:$(go env GOPATH)/bin"
protoc ./pb/*.proto --go_out=./server --go-grpc_out=./ --go_opt=paths=source_relative 
protoc ./pb/*.proto --go_out=./client --go-grpc_out=./ --go_opt=paths=source_relative 
```



