echo "Compiling app with version '$(git describe --tags)'"
go build -o app -ldflags="-X 'main.version=$(git describe --tags)'"