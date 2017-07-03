BIN=horenso-reporter-ciscospark

all: clean test

clean:
	rm -rf build/$(BIN)
	go clean

deps:
	go get -d -t -v ./...
	go get github.com/Songmu/horenso
	go get github.com/jbogarin/go-cisco-spark/ciscospark

test: deps
	go get github.com/stretchr/testify/assert
	go get github.com/pierrre/gotestcover
	go get github.com/mattn/goveralls
	gotestcover -v -covermode=count -coverprofile=.profile.cov ./...
	goveralls -coverprofile=.profile.cov -service=circle-ci -repotoken $${COVERALLS_TOKEN}

deploy_from_circleci: deps
	go get github.com/mitchellh/gox
	go get github.com/tcnksm/ghr
	cd cmd/horenso-reporter-ciscospark/ && \
		gox -osarch "freebsd/amd64 linux/amd64 linux/arm darwin/amd64 windows/amd64" -output "dist/{{.OS}}_{{.Arch}}/{{.Dir}}" && \
		mkdir -p distpkg && \
		for ARCH in `ls dist/`; do zip -j -o distpkg/horenso-reporter-ciscospark_$${ARCH}.zip dist/$${ARCH}/horenso-reporter-ciscospark*; done && \
		ghr -t $${GITHUB_TOKEN} -u $${GITHUB_USER_NAME} -r $${GITHUB_REPO_NAME} -replace $${REPLACE_NAME} distpkg/

.PHONY: all clean deps test deploy_from_circleci
