FROM golang:1.16
WORKDIR /go/src/github.com/lazari24/socket-version-controll
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/lazari24/socket-version-controll/main .
RUN chmod +x main
EXPOSE 8080
CMD ["./main"]