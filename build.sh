

# go install github.com/elazarl/go-bindata-assetfs/...@latest
# go-bindata -o=bindata/bindata.go -pkg=bindata ./cmd/templates/...

version="v0.0.2"
Package="github.com/yzimhao/pinv"
Build_time=`date +%FT%T%z`
COMMIT_SHA1=`git rev-parse HEAD`

CGO_ENABLED=0
go build -ldflags "-X ${Package}.VERSION=${version} -X ${Package}.BUILD=${Build_time} -X ${Package}.COMMIT_SHA1=${COMMIT_SHA1} -X '${Package}.GO_VERSION=`go version`'" -o dist/pinv  cmd/mian.go

cp -rf dist/pinv $GOPATH/bin/pinv