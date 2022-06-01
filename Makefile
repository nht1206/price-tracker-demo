GOFILE = $(shell find . -name '*.go')

default: build

bin:
  mkdir -p bin
  
build: bin/pricetracker

bin/pricetracker: $(GOFILE)
  GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /bin/pricetracker . 
