FROM hub.hamdocker.ir/library/golang:1.20.3

WORKDIR /go/src/github.com/harleywinston/x-manager

COPY ./ .

RUN go build -buildvcs=false -o ./build/manager ./cmd

CMD ["./build/manager"]
