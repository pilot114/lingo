Компиляция:

docker run --rm -v c:/Users/pilot114/sources/go/lingo:/usr/src/myapp -w /usr/src/myapp -e GOOS=windows -e GOARCH=386 golang go build -v

или в интерактивном режиме:
docker run -it --rm -v c:/Users/pilot114/sources/go/lingo:/usr/src/myapp -w /usr/src/myapp golang bash

go get -d -v
go install -v