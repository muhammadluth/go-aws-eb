FROM golang:alpine as build

LABEL maintainer="Muhammad Luthfi <muhammadluthfi059@gmail.com>"

ARG APP_NAME=go-aws-eb

RUN mkdir /app
ADD . /app/

WORKDIR /app
RUN go build -mod=vendor -o ${APP_NAME} .

FROM nginx:alpine
WORKDIR /app
COPY --from=build /app/${APP_NAME}      /app/${APP_NAME}
COPY --from=build /app/${APP_NAME}      /usr/share/nginx/html

EXPOSE 5000
ENTRYPOINT ["/app/go-aws-eb"]
