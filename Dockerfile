# builder image
FROM golang:1.17-alpine as builder
WORKDIR /build
ADD . /build/
RUN apk add build-base --no-cache && go build -a -o feedback-carousell . && rm -rf /usr/local/go/{src,pkg}/*
ENTRYPOINT [ "./feedback-carousell" ]
CMD [ "serve" ]
