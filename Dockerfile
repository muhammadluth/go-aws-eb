FROM golang:alpine as build

LABEL maintainer="Muhammad Luthfi <muhammadluthfi059@gmail.com>"

ARG APP_NAME=go-aws-eb

RUN mkdir /app
ADD . /app/

RUN apk add --upgrade --no-cache tzdata
ENV TZ Asia/Jakarta

WORKDIR /app
RUN go build -mod=vendor -o ${APP_NAME} .

FROM nginx:alpine
WORKDIR /app
COPY --from=build /app/${APP_NAME}      /app/${APP_NAME}
COPY --from=build /app/${APP_NAME}      /usr/share/nginx/html

RUN mkdir -p logs
VOLUME logs

EXPOSE 5000
ENTRYPOINT ["/app/go-aws-eb"]
