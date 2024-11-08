FROM golang:latest as api_builder

WORKDIR /app

COPY api/go.mod api/go.sum ./
RUN go mod download

COPY api/ ./
RUN CGO_ENABLED=0 go build -o server

FROM alpine:3.19

WORKDIR /api

COPY --from=api_builder /app/server /api/server
RUN chmod +x /api/server

CMD [ "./server" ]
