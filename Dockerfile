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
    && apk add cppcheck \
    && apk add ruby \
    && apk add bash

# install dependencies

# Kics
RUN mkdir bin \
    && wget https://github.com/Checkmarx/kics/releases/download/v1.4.9/kics_1.4.9_linux_x64.tar.gz -O "./bin/kics.tar.gz" -q \
    && mkdir ./bin/kics \
    && tar -xf "./bin/kics.tar.gz" -C "./bin/kics" \
    && rm "./bin/kics.tar.gz"

# PMD
RUN wget https://github.com/pmd/pmd/releases/download/pmd_releases%2F6.42.0/pmd-bin-6.42.0.zip -O "./bin/pmd.zip" -q \
    && unzip -q "./bin/pmd.zip" -d "./bin/"  \
    && rm "./bin/pmd.zip"

# gitsecret
RUN git clone https://github.com/awslabs/git-secrets.git ./bin/gitsecret/ \
    &&  cd ./bin/gitsecret/ \
    && make install

# Dependency check
RUN wget https://github.com/jeremylong/DependencyCheck/releases/download/v6.5.0/dependency-check-6.5.0-release.zip -O "./bin/dep_check.zip" -q \
    && unzip -q "./bin/dep_check.zip" -d "./bin" \
    && rm "./bin/dep_check.zip"

RUN gem install bundler-audit
RUN bundle audit update

# cppcheck
RUN mv /usr/bin/cppcheck ./bin/cppcheck

# Dotenv-linter
RUN pip install dotenv-linter

# Keyfinder (IMPORT FAILED)
RUN git clone https://github.com/CERTCC/keyfinder.git ./bin/keyfinder \
    && pip install python-magic
#RUN pip install androguard #INFINITE LOAD
#RUN pip install PyOpenSSL #CRASH

#PyLint
RUN pip install pylint --upgrade

# Build the app
RUN go build -o metascan

# We run the command
# Commenter la ligne en dessous pour que le docker s'arrète pas direct après s'être lancé
ENTRYPOINT ["/go/src/Metascan/metascan"]