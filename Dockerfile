FROM golang:1.24 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-X 'github.com/monakhovm/gcloud-version/cmd.appVersion=2.0.0'" -o /app/app

FROM scratch
COPY --from=builder /app/app .
EXPOSE 8080
LABEL "org.opencontainers.image.source"="https://github.com/monakhovm/gcloud-version"
ENTRYPOINT ["./app"]