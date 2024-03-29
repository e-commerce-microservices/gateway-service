.PHONY: protogen
protogen:
	protoc --proto_path=proto proto/*.proto \
	--go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative

.PHONY: gatewaygen
gatewaygen:
	protoc --proto_path=proto proto/*.proto \
	--go_out=pb --go_opt=paths=source_relative \
	--grpc-gateway_out pb \
	--grpc-gateway_opt paths=source_relative \
	--grpc-gateway_opt grpc_api_configuration=route.yaml

.PHONY: rebuild
rebuild:
	docker build -t ngoctd/ecommerce-gateway:latest . && \
	docker push ngoctd/ecommerce-gateway

.PHONY: redeploy
redeploy:
	kubectl rollout restart deployment depl-gateway
	