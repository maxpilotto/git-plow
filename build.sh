platforms=("darwin;amd64" "linux;amd64" "windows;amd64")
version=$(git describe --tags --abbrev=0)

for platform in ${platforms[@]}; do
    os=${platform%;*}
    arch=${platform#*;}
    env GOOS=$os GOARCH=$arch go build -ldflags "-X main.Version=$version" -o dist/git-plow.$os.$arch cmd/git-plow.go
done