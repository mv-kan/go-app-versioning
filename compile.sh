echo 
go build -o app -ldflags="-X 'main.version=$(git describe --tags)'"