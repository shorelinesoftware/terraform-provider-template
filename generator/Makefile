ifeq ($(USE_LOCAL), 1)
	include ./env/dev/local.env
	export
else ifeq ($(PROVIDER_BRAND), shoreline)
	include ./env/release/shoreline.env
	export
else ifeq ($(PROVIDER_BRAND), ngg)
	include ./env/release/ngg.env
	export
endif

default: install

REPODIR=/tmp/tf-repo/providers

BINARY=terraform-provider-$(PROVIDER_SHORT_NAME)
VERSION=1.15.30

BUILD_ENV_VARS=-ldflags "-X 'main.RenderedProviderName=\"$(RENDERED_PROVIDER_NAME)\"' -X 'main.ProviderShortName=\"$(PROVIDER_SHORT_NAME)\"' -X 'provider.RenderedProviderName=\"$(RENDERED_PROVIDER_NAME)\"' -X 'provider.ProviderShortName=\"$(PROVIDER_SHORT_NAME)\"' -X 'provider.EnvVarsNamePrefix=\"$(ENV_VARS_NAME_PREFIX)\"' -X 'provider.TfLogFile=\"$(TF_LOG_FILE)\"' -X 'provider.DefaultUserName=\"$(DEFAULT_USER_NAME)\"'"


// NOTE: this only works for 64 bit linux and MacOs ("darwin")
OS=$(shell uname | tr 'A-Z' 'a-z')
SUBPATH=$(PROVIDER_PATH)/local/$(PROVIDER_SHORT_NAME)/$(VERSION)/$(OS)_amd64


.PHONY: generate
generate:
	go run ./*.go


build: generate format
	go generate $(BUILD_ENV_VARS)
	go build $(BUILD_ENV_VARS) -o ./$(BINARY)

test:
	echo unit-tests...

check:
	gofmt -l .

format:
	gofmt -w .

# NOTE: This relies on your ~/.terraformrc pointing to /tmp/tf-repo.
#   See terraformrc in the current dir
install: build
	rm -rf $(REPODIR)/*
	mkdir -p $(REPODIR)/$(SUBPATH)
	cp $(BINARY) $(REPODIR)/$(SUBPATH)/$(BINARY)

# This sets up your ~/.terraformrc (NOTE: need to re-run when the version changes)
use_local: 
	@echo 'Setting up local overrides for terraform provider in ~/.terraformrc'
	@echo 'NOTE: You need to re-run "make use_local" when the version changes."'
	@echo 'provider_installation { dev_overrides { "$(PROVIDER_PATH)" = "$(REPODIR)/$(SUBPATH)" } }' > ${HOME}/.terraformrc
