FROM golang:latest AS build
WORKDIR /go/src/goproxy
COPY . .
RUN go install

FROM buildpack-deps:buster-scm
COPY --from=build /go/bin/goproxy /usr/local/bin/
VOLUME /var/db/goproxy
EXPOSE 80
CMD exec goproxy
