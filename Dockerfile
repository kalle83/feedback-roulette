# builder image
FROM golang:1.17-alpine as builder
WORKDIR /build
ADD . /build/
RUN apk add npm build-base --no-cache \
  && go build -a -o feedback-carousell . \
  && rm -rf /usr/local/go/{src,pkg}/* \
  && npm install --prefix ui/ \
  && npm run build --prefix ui/
ENTRYPOINT [ "./feedback-carousell" ]
CMD [ "serve" ]
