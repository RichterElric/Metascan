FROM golang:1.16-alpine as Metascan

WORKDIR /go/src/Metascan
COPY ./Metascan ./

# Install base tools
RUN apk update \
    && apk upgrade \
    && apk add python3 \
    && apk add py3-pip \
    && apk add wget \
    && apk add git \
    && apk add libc6-compat \
    && apk add zip \
    && apk add make \
    && apk add yarn \
    && apk add openjdk11 \
    && apk add bash

# install dependencies

# Kics
RUN mkdir bin \
    && wget https://github.com/Checkmarx/kics/releases/download/v1.4.8/kics_1.4.8_linux_x64.tar.gz -O "./bin/kics.tar.gz" -q \
    && mkdir ./bin/kics \
    && tar -xf "./bin/kics.tar.gz" -C "./bin/kics" \
    && rm "./bin/kics.tar.gz"

# gitsecret
RUN git clone https://github.com/awslabs/git-secrets.git ./bin/gitsecret/ \
    &&  cd ./bin/gitsecret/ \
    && make install

# Dependency check
RUN wget https://github.com/jeremylong/DependencyCheck/releases/download/v6.5.0/dependency-check-6.5.0-release.zip -O "./bin/dep_check.zip" -q \
    && unzip -q "./bin/dep_check.zip" -d "./bin" \
    && rm "./bin/dep_check.zip"

# Build the app
RUN go build -o metascan

# We run the command
# Commenter la ligne en dessous pour que le docker s'arrète pas direct après s'être lancé
ENTRYPOINT ["/go/src/Metascan/metascan"]