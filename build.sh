# Quick build script for building for multiple platforms

# mac build
echo "Building darwin/amd64"
mkdir -p dist/darwin/amd64
GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build
cp README.md  dist/darwin/amd64
cp cwl dist/mac/

# linux build
echo "Building linux/amd64"
mkdir -p dist/linux/amd64
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build
cp README.md  dist/darwin/amd64
cp cwl dist/linux/amd64/

