OUTPUT = release/ft
## linux: 编译打包linux
.PHONY: linux win mac

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${OUTPUT} .

## win: 编译打包win
win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${OUTPUT}.exe .

## mac: 编译打包mac
mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ${OUTPUT}_mac .

## mac_arm: 编译打包mac_arm
mac_arm:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o ${OUTPUT}_mac_arm .

release: linux mac win mac_arm