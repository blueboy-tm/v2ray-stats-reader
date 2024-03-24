GOOS=windows GOARCH=amd64 go build -o bin/v2ray-stats-reader-win-amd64.exe main.go
GOOS=windows GOARCH=arm64 go build -o bin/v2ray-stats-reader-win-arm64.exe main.go
GOOS=darwin GOARCH=arm64 go build -o bin/v2ray-stats-reader-darwin-arm64 main.go
GOOS=darwin GOARCH=amd64 go build -o bin/v2ray-stats-reader-darwin-amd64 main.go
GOOS=linux GOARCH=amd64 go build -o bin/v2ray-stats-reader-linux-amd64 main.go
GOOS=linux GOARCH=386 go build -o bin/v2ray-stats-reader-linux-i386 main.go
