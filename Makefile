PREFIX_PATH := ${HOME}/.local
TEMP_PATH := /tmp/packets
CWD_PATH = $(shell pwd)

deps: deps_structure deps_grpc deps_ui deps_server

deps_structure:
	mkdir -p ${PREFIX_PATH}
	mkdir -p ${TEMP_PATH}

deps_grpc:
	if [ ! command -v protoc &> /dev/null ]; then \
		if [ ! -d "${TEMP_PATH}/grpc" ]; then \
			git clone \
				--recurse-submodules \
				-b v1.62.0 \
				--depth 1 \
				--shallow-submodules https://github.com/grpc/grpc \
				${TEMP_PATH}/grpc; \
		fi; \
		cd ${TEMP_PATH}/grpc/ && \
			mkdir -p ${TEMP_PATH}/grpc/cmake/build/ && \
			pushd ${TEMP_PATH}/grpc/cmake/build/ && \
			cmake -DgRPC_INSTALL=ON \
			-DgRPC_BUILD_TESTS=OFF \
			-DCMAKE_INSTALL_PREFIX=${PREFIX_PATH} \
			../.. && \
			make -j 4 && \
			make install; \
	fi


deps_ui:
	dart pub global activate protoc_plugin

deps_server:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

build: build_server

build_deps: deps

build_server: protoc_server
	cd ${CWD_PATH}/server && \
		go mod tidy && \
		go build -o ${CWD_PATH}/target/server

protoc: protoc_ui protoc_server

protoc_ui:
	protoc \
		--dart_out=grpc:${CWD_PATH}/ui/lib/api \
		--plugin=protoc-gen-dart=${HOME}/.pub-cache/bin/protoc-gen-dart \
		protos/api.proto

protoc_server:
	protoc \
		--go_out=${CWD_PATH}/server/ \
		--go-grpc_out=${CWD_PATH}/server/ \
		--plugin=protoc-gen-go=$(shell go env GOPATH)/bin/protoc-gen-go \
		--plugin=protoc-gen-go-grpc=$(shell go env GOPATH)/bin/protoc-gen-go-grpc \
		protos/api.proto

clean:
	rm -rf ${CWD_PATH}/target
	rm -rf ${TEMP_PATH}