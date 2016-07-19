execbeat:
	go build

.PHONY: getdeps
getdeps:
	glide up --strip-vcs --update-vendored

.PHONY: test
test:
	go test . ./beat/...

.PHONY: updatedeps
updatedeps:
	glide up --strip-vcs --update-vendored

.PHONY: install_cfg
install_cfg:
	cp etc/execbeat.yml $(PREFIX)/execbeat-linux.yml
	cp etc/execbeat.template.json $(PREFIX)/execbeat.template.json
	# darwin
	cp etc/execbeat.yml $(PREFIX)/execbeat-darwin.yml
	# win
	cp etc/execbeat.yml $(PREFIX)/execbeat-win.yml

.PHONY: gofmt
gofmt:
	go fmt ./...

.PHONY: cover
cover:
	echo 'mode: atomic' > coverage.txt && go list . ./beat | xargs -n1 -I{} sh -c 'go test -covermode=atomic -coverprofile=coverage.tmp {} && tail -n +2 coverage.tmp >> coverage.txt' && rm coverage.tmp

.PHONY: clean
clean:
	rm -r cover || true
	rm profile.cov || true
	rm execbeat || true
	rm coverage.txt || true
