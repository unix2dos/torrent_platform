.SILENT :
.PHONY : clean

NAME := torrent_platform
PRE := ks
ROOF := $(NAME)


all: clean dht mac download


clean:
	echo "Cleaning dist"
	rm -rf dist fe/build
	rm -f $(NAME) $(NAME)-*
	rm -f ./ks-*


dht:
	echo "Building $@"
	mkdir -p dist/linux_amd64 && GOOS=linux GOARCH=amd64 go build -o dist/linux_amd64/$(PRE)-$@ $(ROOF)/cmd/$(PRE)-$@
	#mkdir -p dist/darwin_amd64 && GOOS=darwin GOARCH=amd64 go build -o dist/darwin_amd64/$(PRE)-$@ $(ROOF)/cmd/$(PRE)-$@
.PHONY: dht

mac:
	echo "Building $@"
	mkdir -p dist/linux_amd64 && GOOS=linux GOARCH=amd64 go build -o dist/linux_amd64/$(PRE)-$@ $(ROOF)/cmd/$(PRE)-$@
	#mkdir -p dist/darwin_amd64 && GOOS=darwin GOARCH=amd64 go build  -o dist/darwin_amd64/$(PRE)-$@ $(ROOF)/cmd/$(PRE)-$@
.PHONY: mac


download:
	echo "Building $@"
	mkdir -p dist/linux_amd64 && GOOS=linux GOARCH=amd64 go build  -o dist/linux_amd64/$(PRE)-$@ $(ROOF)/cmd/$(PRE)-$@
	#mkdir -p dist/darwin_amd64 && GOOS=darwin GOARCH=amd64 go build -o dist/darwin_amd64/$(PRE)-$@ $(ROOF)/cmd/$(PRE)-$@
.PHONY: download
