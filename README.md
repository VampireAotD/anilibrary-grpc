# Anilibrary gRPC

Microservice that generates all gRPC clients/servers for Anilibrary

[![tests](https://github.com/VampireAotD/anilibrary-grpc/actions/workflows/tests.yml/badge.svg)](https://github.com/VampireAotD/anilibrary-grpc/actions/workflows/tests.yml)

---

## Issues

Currently, gRPC client for [Anilibrary](https://github.com/VampireAotD/anilibrary) cannot be generated properly, because
of [validate plugin](https://github.com/bufbuild/protoc-gen-validate) that uses `proto2 syntax`, which is [not supported
by PHP](https://github.com/protocolbuffers/protobuf/issues/3623) and cannot
be [excluded from generation with buf](https://github.com/bufbuild/buf/issues/645), so
right now [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) is being used.

---

## Buf

For simplicity, [buf](https://github.com/bufbuild/buf) is used to generate gRPC clients and servers.

You can install it [locally](https://docs.buf.build/installation) or
use [docker image](https://hub.docker.com/r/bufbuild/buf).

---

## Generate

Using docker image

```shell
make generate
```

Using local binary

```shell
buf generate
```