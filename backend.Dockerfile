FROM golang:1.24.1-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /go
COPY ./backend .

RUN go install
RUN go build -o /bin/backend

FROM scratch
COPY --from=builder /go/bin/backend /go/bin/backend
ARG VITE_BACKEND_PORT
EXPOSE ${VITE_BACKEND_PORT}
ENTRYPOINT ["/go/bin/backend"]
