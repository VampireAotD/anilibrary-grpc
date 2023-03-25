buf := docker run --rm --volume "$(shell pwd):/buf" --workdir /buf bufbuild/buf

.PHONY: generate
generate:
	$(buf) generate --exclude-path proto/scraper/specification

.PHONY: generate-specification
generate-specification:
	$(buf) generate --template buf.gen-spec.yml

.PHONY: lint
lint:
	$(buf) lint

.PHONY: build
build:
	$(buf) build

.PHONY: mod-update
mod-update:
	$(buf) mod update