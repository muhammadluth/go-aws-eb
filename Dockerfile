FROM golang:alpine as build

LABEL maintainer="Muhammad Luthfi <muhammadluthfi059@gmail.com>"

ARG APP_NAME=go-aws-eb

RUN mkdir /app
ADD . /app/

WORKDIR /app
RUN go build -mod=vendor -o ${APP_NAME} .

FROM alpine
WORKDIR /app
COPY --from=build /app/${APP_NAME}  /app/${APP_NAME}
EXPOSE 5000

ENTRYPOINT ["/app/go-aws-eb"]
