/*
 * @Description  : Makefile
 * @Author       : jsmjsm
 * @Github       : https://github.com/jsmjsm
 * @Date         : 2021-03-08 19:29:48
 * @LastEditors  : jsmjsm
 * @LastEditTime : 2021-03-08 19:30:06
 * @FilePath     : /goto/goto_v4/Makefile
 */

BINARY="goto_v4"
VERSION=1.0.3
BUILD=`date +%FT%T%z`

PACKAGES=`go list ./... | grep -v /vendor/`
VETPACKAGES=`go list ./... | grep -v /vendor/ | grep -v /examples/`
GOFILES=`find . -name "*.go" -type f -not -path "./vendor/*"`

default:
	@go build -o ${BINARY} -tags=${VERSION}

list:
	@echo ${PROJECT}
	@echo ${PACKAGES}
	@echo ${VETPACKAGES}
	@echo ${GOFILES}

fmt:
	@gofmt -s -w ${GOFILES}

install:
	@govendor sync -v

test:
	@go test -cpu=1,2,4 -v -tags integration ./...

vet:
	@go vet $(VETPACKAGES)

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

run:
	@go build -o ${BINARY} -tags=${VERSION} && ./${BINARY}

.PHONY: default fmt install test vet clean run

