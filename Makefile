all: get_deps install doc

GO=$(shell which go)
GOGET=$(GO) get

VERSION=v0.0.1
DESCRIPTION=version 0.0.1

clean:
	rm -rf dist

get_deps:
	@echo -n "get dependencies... "
	#@$(GOGET) github.com/moovweb/gokogiri
	@echo ok

install:
	go install

gorelease:
	-git tag -d $(VERSION)
	git tag -a $(VERSION) -m "$(DESCRIPTION)"
	goreleaser

.PHONY: install
