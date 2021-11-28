# builder image
FROM golang:1.17-alpine as builder
RUN apk add build-base 
WORKDIR /build
ADD . /build/
RUN go build -a -o feedback-carousell .


# generate clean, final image for end users
FROM alpine
COPY --from=builder /build/feedback-carousell .
ADD ui/dist ui/dist
ADD questions.json .

# executable
ENTRYPOINT [ "./feedback-carousell" ]
# arguments that can be overridden
CMD [ "serve" ]
