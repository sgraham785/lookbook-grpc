FROM golang:1.9.1 AS dev-env

# ARG USER
# ARG UID

# RUN useradd $USER -o -u $UID
# USER $USER

WORKDIR /go/src/app
