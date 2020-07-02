FROM golang:1.14 as dev
WORKDIR /go/src/github.com/whuangz/product-api
RUN GO111MODULE=on go get github.com/cortesi/modd/cmd/modd
COPY . .
CMD modd

FROM golang:1.14 as build
WORKDIR /product-api
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app .

FROM gcr.io/distroless/static as prod
COPY --from=build /product-api/app /app
CMD ["/app"]