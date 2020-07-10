FROM golang:1.14 as dev
WORKDIR /blog-api
ENV GO111MODULE=on 
RUN go get github.com/cortesi/modd/cmd/modd
COPY go.mod .
COPY go.sum .
COPY . .
CMD modd

FROM golang:1.14 as build
WORKDIR /blog-api
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app .

FROM gcr.io/distroless/static as prod
COPY --from=build blog-api/app /app
CMD ["/app"]