# FROM golang:alpine AS build-env
FROM golang:alpine
LABEL maintainer "francoishill11@gmail.com"

RUN echo Installing git \
  apk update && apk upgrade \
  && apk add --no-cache bash git openssh

ENV CGO_ENABLED=0

COPY . $GOPATH/src/github.com/golang-devops/gen/

RUN echo Installing gen with all its TypeWriters \
  && cd $GOPATH/src/github.com/golang-devops/gen && go install \
  && gen add github.com/clipperhouse/slice \
  && gen add github.com/clipperhouse/set \
  && gen add github.com/clipperhouse/stringer \
  && gen add github.com/nickmab/gen/typewriters/container \
  && gen add github.com/clipperhouse/linkedlist \
  && gen add github.com/ninibe/atomicmapper/gen \
  && gen add github.com/clipperhouse/ring \
  && gen add github.com/clipperhouse/set \
  && gen add github.com/jackc/signal \
  && gen add github.com/svett/gen/stack \
  && gen add github.com/michaelsmanley/flags \
  && gen add github.com/fernandokm/sequences

# FROM alpine
# COPY --from=build-env /go/bin/gen /usr/bin/
# COPY --from=build-env /root/_user_gen.go /root/
# RUN chmod +x /usr/bin/gen

RUN cp /go/bin/gen /usr/bin/ \
  && chmod +x /usr/bin/gen

ENTRYPOINT ["/usr/bin/gen"]