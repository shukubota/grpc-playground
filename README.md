# grpc-playground
## 概要
gRPCのあれこれ
- grpc
- grpc-gateway
- grpc-web
- connect

## 動作
### grpc-gateway
#### コンパイル
```shell
buf generate
```
でgen配下にコードが生成される

#### サーバ
```shell
go run main.go
```
でgrpcサーバ(port: 5001)とgrpc-gatewayのプロキシ(port: 8085)がたつ

```shell
curl -X GET http://localhost:8085/example-messages
```

で
```json
{"message":"OK"}
```
が返ってくる。
