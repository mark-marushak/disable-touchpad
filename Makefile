
.PHONY: commit-push
commit-push:
	go build -o bin/disable-touchpad cmd/disable-touchpad/disable-touchpad.go
	git add . ;
	git commit -m "$(message)"
	git push origin "$(branch)"

.PHONY: tag
tag:
	git tag "$1"
	git push origin "$1"


.PHONY: build
build:
	go build -o bin/disable-touchpad cmd/disable-touchpad/disable-touchpad.go
