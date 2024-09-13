# base image for dragoscojocaru/forxy
FROM golang:1.22-alpine3.20 AS build

WORKDIR /go/src/forxy

COPY . /go/src/forxy

RUN go mod download && go build

FROM alpine:3.20

ENV FORXY_CONFIG_PATH=/etc/forxy/forxy.yaml

COPY --from=build /go/src/forxy/forxy /etc/forxy/forxy
COPY --from=build /go/src/forxy/forxy.yaml /etc/forxy/forxy.yaml

CMD ["/etc/forxy/forxy"]