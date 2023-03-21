
start:
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sudo sh -s -- -b $(go env GOPATH)/bin; \
	alias air='~/.air'; \
	air -c .air.toml

build-air:
	swag init
	go build -o ./tmp/main .