FROM golang:1.22 as builder

WORKDIR /app

ADD go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o bookstore ./cmd

FROM scratch

COPY --from=builder /app/bookstore /
COPY config.yml ./config.yml

ENTRYPOINT ["./bookstore"]


