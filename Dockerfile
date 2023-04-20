FROM golang:1.19-alpine3.16 AS build
LABEL MAINTAINER = 'Auth (dany0814)'
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /vakalan-backend -ldflags="-w -s" ./cmd/main.go

FROM scratch
WORKDIR /
COPY --from=build vakalan-backend/ /vakalan-backend

CMD ["/vakalan-backend"]