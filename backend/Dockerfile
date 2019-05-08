FROM golang:1.12 as builder

WORKDIR /go/src/github.com/canobbioe/reelo
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/reelo .


FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/bin/reelo .

