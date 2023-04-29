#!/usr/bin/env bash
#
# Generate all protobuf bindings.
# Run from repository root.
set -e
set -u

if ! [[ "$0" =~ "scripts/genproto.sh" ]]; then
	echo "must be run from repository root"
	exit 255
fi

if ! [[ $(protoc --version) =~ "3.6.1" ]]; then
	echo "could not find protoc 3.6.1, is it installed + in PATH?"
	exit 255
fi

# Since we run go install, go mod download, the go.sum will change.
# Make a backup.
cp go.sum go.sum.bak

INSTALL_PKGS=""
for pkg in ${INSTALL_PKGS}; do
    GO111MODULE=on go install "$pkg"
done

RELEASENAMER_ROOT="${PWD}"
RELEASENAMER_PATH="${RELEASENAMER_ROOT}/releasenamer"
GOGOPROTO_ROOT="$(GO111MODULE=on go list -mod=readonly -f '{{ .Dir }}' -m github.com/gogo/protobuf)"
GOGOPROTO_PATH="${GOGOPROTO_ROOT}:${GOGOPROTO_ROOT}/protobuf"
GRPC_GATEWAY_ROOT="$(GO111MODULE=on go list -mod=readonly -f '{{ .Dir }}' -m github.com/grpc-ecosystem/grpc-gateway)"

DIRS="releasenamer"

echo "generating code"
for dir in ${DIRS}; do
	pushd ${dir}
	  echo $PWD
		protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative release_namer.proto

#		sed -i.bak -E 's/import _ \"google\/protobuf\"//g' -- *.pb.go
#		sed -i.bak -E 's/\t_ \"google\/protobuf\"//g' -- *.pb.go
#		sed -i.bak -E 's/golang\/protobuf/gogo\/protobuf/g' -- *.go
		rm -f -- *.bak
	popd
done

mv go.sum.bak go.sum
