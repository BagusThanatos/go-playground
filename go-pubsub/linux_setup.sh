#PlEASE HAVE GO INSTALLED BEFOREHAND
# set GOPATH=./
# set GOBIN=./build/

echo "RUN TEST"
go test ./test/backend/... -v

echo "BUILDING"
mkdir -p build/linux/
go build -o build/linux/main  ./src/backend/...
./build/linux/main
