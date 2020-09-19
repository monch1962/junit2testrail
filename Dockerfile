# build stage
FROM golang:alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go build -o junit2testrail

# final stage
FROM scratch
COPY --from=builder /build/junit2testrail /app/
WORKDIR /app
CMD ["./junit2testrail"]