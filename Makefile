all: truck

truck:
	go generate
	go build -o truck

cross: darwin.tgz freebsd-386.tgz freebsd-amd64.tgz linux-386.tgz linux-amd64.tgz

darwin.tgz:
	GOOS=darwin GOARCH=amd64 go build -o truck
	tar -cvzf $@ truck

freebsd-386.tgz:
	GOOS=freebsd GOARCH=386 go build -o truck
	tar -cvzf $@ truck

freebsd-amd64.tgz:
	GOOS=freebsd GOARCH=amd64 go build -o truck
	tar -cvzf $@ truck

linux-386.tgz:
	GOOS=linux GOARCH=386 go build -o truck
	tar -cvzf $@ truck

linux-amd64.tgz:
	GOOS=linux GOARCH=amd64 go build -o truck
	tar -cvzf $@ truck
