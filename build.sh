env GOOS=$os GOARCH=$arch go build -ldflags "-X main.Version=$(git describe --tags --abbrev=0)" -o dist/git-plow git-plow.go
