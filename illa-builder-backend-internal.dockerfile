# ---------------------
# build illa-builder-backend-internal
FROM golang:1.19-bullseye as builder-for-backend

## set env
ENV LANG C.UTF-8
ENV LC_ALL C.UTF-8

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

## build
WORKDIR /opt/illa/illa-builder-backend
RUN cd  /opt/illa/illa-builder-backend
RUN ls -alh

COPY ./ ./

RUN cat ./Makefile

RUN make build-http-server-internal

RUN ls -alh ./bin/illa-builder-backend-internal


# -------------------
# build runner images
FROM alpine:latest as runner

WORKDIR /opt/illa/illa-builder-backend/bin/

## copy backend bin
COPY --from=builder-for-backend /opt/illa/illa-builder-backend/bin/illa-builder-backend-internal /opt/illa/illa-builder-backend/bin/


RUN ls -alh /opt/illa/illa-builder-backend/bin/



# run
EXPOSE 9005
CMD ["/bin/sh", "-c", "/opt/illa/illa-builder-backend/bin/illa-builder-backend-internal"]
