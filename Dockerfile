FROM golang:1.15-alpine AS GO_BUILD
RUN apk add git
COPY . /webglobalinsurancedocker
WORKDIR /webglobalinsurancedocker
RUN go build -o /go/bin/webglobalinsurancedocker
FROM alpine:3.10
WORKDIR app
COPY --from=GO_BUILD /go/bin/webglobalinsurancedocker/ ./
EXPOSE 7070
CMD ./webglobalinsurancedocker