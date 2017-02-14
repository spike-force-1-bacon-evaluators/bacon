FROM golang:1.7.4-wheezy

COPY . src/github.com/github.com/bacon/bacon

WORKDIR src/github.com/github.com/bacon/bacon

RUN go build

ENTRYPOINT ["./bacon"]
