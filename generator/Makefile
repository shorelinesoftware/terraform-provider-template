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

BUILD_ENV_VARS=-ldflags "-X 'main.RenderedProviderName=\"$(RENDERED_PROVIDER_NAME)\"' -X 'main.ProviderShortName=\"$(PROVIDER_SHORT_NAME)\"' -X 'provider.RenderedProviderName=\"$(RENDERED_PROVIDER_NAME)\"' -X 'provider.ProviderShortName=\"$(PROVIDER_SHORT_NAME)\"' -X 'provider.EnvVarsNamePrefix=\"$(ENV_VARS_NAME_PREFIX)\"' -X 'provider.TfLogFile=\"$(TF_LOG_FILE)\"' -X 'provider.DefaultUserName=\"$(DEFAULT_USER_NAME)\"'"

.PHONY: generate
generate:
	go run ./*.go


build: generate format
	go generate $(BUILD_ENV_VARS)
	go build $(BUILD_ENV_VARS)

check:
	gofmt -l .

format:
	gofmt -w .
