default: build
all: package

export GOPATH=$(CURDIR)/
export GOBIN=$(CURDIR)/.temp/

init: clean
	go get ./...

build: init fmt
	go build -o ./.output/etymu .

test:
	go test
	go test -bench=.

clean:
	@rm -rf ./.output/

fmt:
	@go fmt .
	@go fmt ./src/etymu

sample: build
	@./.output/etymu -l go -o ./.temp/calc ./sample/calc.el

dist: build test

	export GOOS=linux; \
	export GOARCH=amd64; \
	go build -o ./.output/etymu64 .

	export GOOS=linux; \
	export GOARCH=386; \
	go build -o ./.output/etymu32 .

	export GOOS=darwin; \
	export GOARCH=amd64; \
	go build -o ./.output/etymu_osx .

	export GOOS=windows; \
	export GOARCH=amd64; \
	go build -o ./.output/etymu.exe .

package: versionTest fpmTest dist

	fpm \
		--log error \
		-s dir \
		-t deb \
		-v $(ETYMU_VERSION) \
		-n etymu \
		./.output/etymu64=/usr/local/bin/etymu \
		./docs/etymu.7=/usr/share/man/man7/etymu.7 \
		./autocomplete/etymu=/etc/bash_completion.d/etymu

	fpm \
		--log error \
		-s dir \
		-t deb \
		-v $(ETYMU_VERSION) \
		-n etymu \
		-a i686 \
		./.output/etymu32=/usr/local/bin/etymu \
		./docs/etymu.7=/usr/share/man/man7/etymu.7 \
		./autocomplete/etymu=/etc/bash_completion.d/etymu

	@mv ./*.deb ./.output/

	fpm \
		--log error \
		-s dir \
		-t rpm \
		-v $(ETYMU_VERSION) \
		-n etymu \
		./.output/etymu64=/usr/local/bin/etymu \
		./docs/etymu.7=/usr/share/man/man7/etymu.7 \
		./autocomplete/etymu=/etc/bash_completion.d/etymu
	fpm \
		--log error \
		-s dir \
		-t rpm \
		-v $(ETYMU_VERSION) \
		-n etymu \
		-a i686 \
		./.output/etymu32=/usr/local/bin/etymu \
		./docs/etymu.7=/usr/share/man/man7/etymu.7 \
		./autocomplete/etymu=/etc/bash_completion.d/etymu

	@mv ./*.rpm ./.output/

fpmTest:
ifeq ($(shell which fpm), )
	@echo "FPM is not installed, no packages will be made."
	@echo "https://github.com/jordansissel/fpm"
	@exit 1
endif

versionTest:
ifeq ($(ETYMU_VERSION), )

	@echo "No 'ETYMU_VERSION' was specified."
	@echo "Export a 'ETYMU_VERSION' environment variable to perform a package"
	@exit 1
endif
