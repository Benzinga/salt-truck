all: truck

truck:
	go generate
	go build -o truck
