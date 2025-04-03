FROM golang:1.24.1-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /go
COPY ./backend .

RUN go install
RUN go build -o /bin/backend

FROM scratch
COPY --from=builder /go/bin/backend /go/bin/backend
ENTRYPOINT ["/go/bin/backend"]
