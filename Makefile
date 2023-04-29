# Needs to be defined before including Makefile.common to auto-generate targets
DOCKER_ARCHS ?= amd64 armv7 arm64 ppc64le s390x

GOLANGCI_LINT_OPTS ?= --timeout 4m

include Makefile.common

DOCKER_IMAGE_NAME       ?= release-namer

.PHONY: test
# If we only want to only test go code we have to change the test target
# which is called by all.
ifeq ($(GO_ONLY),1)
test: common-test
else
test: common-test ui-build-module ui-test ui-lint
endif

.PHONY: tarball
tarball: common-tarball

.PHONY: docker
docker: common-docker

plugins/plugins.go: plugins.yml plugins/generate.go
	@echo ">> creating plugins list"
	$(GO) generate -tags plugins ./plugins

.PHONY: plugins
plugins: plugins/plugins.go

.PHONY: build
build: common-build
