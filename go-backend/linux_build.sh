#PlEASE HAVE GO INSTALLED BEFOREHAND
set GOPATH=./
set GOBIN=./build/
echo "RUN TEST"
go test ./test/backend/... -v

go build -o build/linux/main  ./src/backend/...
./build/linux/main
