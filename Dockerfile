FROM golang:1.14 as dev
WORKDIR /go/src/github.com/whuangz/microservice
RUN GO111MODULE=on go get github.com/cortesi/modd/cmd/modd
COPY . .
CMD modd

FROM golang:1.14 as build
WORKDIR /microservice
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app .

FROM gcr.io/distroless/static as prod
COPY --from=build /microservice/app /app
CMD ["/app"]