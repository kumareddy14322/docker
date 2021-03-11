FROM golang:1.14-alpine AS build
ENV SRC_DIR=/go/src/github.com/kumareddy14322/
ENV GOBIN=/go/bin
WORKDIR $GOBIN
# Add the source code:
ADD . $SRC_DIR
RUN cd /go/src/;
RUN go install github.com/kumareddy14322/;
ENTRYPOINT ["./kumar"]
EXPOSE 8080