FROM golang:1.19.3 as builder
COPY ${pwd} /app
WORKDIR /app
RUN CGO_ENABLED=1 go build -ldflags '-s -w -extldflags "-static"' -o /app/appbin cmd/main.go

FROM gcr.io/distroless/base-debian11
LABEL MAINTAINER = "Kamran Karimov"
WORKDIR /app
COPY --from=builder /app /app
ENV GIN_MODE release
EXPOSE 8080 8081

CMD ["./appbin"]