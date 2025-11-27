####
# 参数:
# --config-path: 指定配置文件的路径，默认在./config文件夹下面
#
####
export PATH := $(shell go env GOPATH)/bin:$(PATH)


### 本地使用order
.PHONY: run-local-order
run-local-order:
	go run ./internal/order/cmd
.PHONY: run-local-stock
run-local-stock:
	go run ./internal/stock/cmd



### 生成代码的script
.PHONY: generate-protoc
generate-protoc:
	protoc \
		--proto_path=./api/proto \
		--go_out=./internal/common/genproto/orderpb \
		--go_opt=module=github.com/looksaw2/gorder3/internal/common/genproto/orderpb \
		--go-grpc_out=./internal/common/genproto/orderpb \
		--go-grpc_opt=module=github.com/looksaw2/gorder3/internal/common/genproto/orderpb \
		order.proto

.PHONY: generate-oapi
generate-oapi:
	oapi-codegen \
		-package  order \
		-generate "types" \
		-o ./internal/common/client/order/types.gen.go ./api/openapi/order.yaml
	oapi-codegen \
		-package order \
		-generate "chi-server" \
		-o ./internal/common/client/order/server.gen.go ./api/openapi/order.yaml
	oapi-codegen \
		-package order \
		-generate "spec" \
		-o ./internal/common/client/order/spec.gen.go ./api/openapi/order.yaml



###函数测试使用
test-order-get:
	curl -X GET localhost:8282/api/customer/1/orders/2
test-order-post:
	curl -X POST localhost:8282/api/customer/22/orders