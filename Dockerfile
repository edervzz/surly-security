# syntax=docker/dockerfile:1

# docker image name:tag
FROM golang:1.20.6-alpine3.18 AS builder
# set folder
WORKDIR /source/
# copy *.go files to /app
COPY /cmd/. ./cmd/
COPY /docs/. ./docs/
COPY /internal/. ./internal/
COPY /toolkit/. ./toolkit/
COPY go.mod go.sum ./
# run command to download golang project dependencies
RUN go mod download
# go build
RUN CGO_ENABLED=0 GOOS=linux go build -o ./build/surly-security ./cmd/main.go

# final stage
FROM alpine
WORKDIR /app/
COPY --from=builder /source/build/. ./
# default envars
ENV BASE_PATH=/security/api/v1
ENV Jwt:Issuer=surly-security.com
ENV Jwt:Audience=surly-eco.com
ENV Jwt:Expiration=1800
ENV Jwt:RefreshExpiration=10080
ENV Jwt:Key=
ENV APP_PORT=
ENV DB_SERVER=
ENV DB_NAME=
ENV DB_PORT=
ENV DB_USR=
ENV DB_PWD=

CMD ["/app/surly-security"]









# ENV APP_PORT=6001
# ENV DB_SERVER=localhost
# ENV DB_NAME=surly_security
# ENV DB_PORT=3306
# ENV DB_USR=root
# ENV DB_PWD=eder
# ENV Jwt:Key=S3cr3tP@ss!S3cr3tP@ss!S3cr3tP@ss!S3cr3tP@ss!S3cr3tP@ss!S3cr3tP@ss!
