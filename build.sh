# OS별로 빌드함.
GOOS=linux GOARCH=amd64 go build -o ./bin/linux/download download.go
GOOS=windows GOARCH=amd64 go build -o ./bin/windows/download.exe download.go
GOOS=darwin GOARCH=amd64 go build -o ./bin/darwin/download download.go

# Github Release에 업로드 하기위해 압축
cd ./bin/linux/ && tar -zcvf ../download_linux.tgz . && cd -
cd ./bin/windows/ && tar -zcvf ../download_windows.tgz . && cd -
cd ./bin/darwin/ && tar -zcvf ../download_darwin.tgz . && cd -
