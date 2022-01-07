#!/bin/bash

SCRIPT_DIR="$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

if [ "$EUID" -ne 0 ]
then
  echo "Please run as root"
  exit
fi

echo $SCRIPT_DIR
### Install utilities ###
# Check if go is installed
if ! command -v go &> /dev/null
then
  echo "go not found, downloading"
  apt-get install go
fi
# Check if python is installed
if ! command -v python3 &> /dev/null
then
  echo "python3 not found, downloading"
  apt-get install python3
fi
# Installing pip
if ! command -v pip3 &> /dev/null
then
  echo "pip3 not found, downloading"
  apt-get install pip3
fi
# tar not found
if ! command -v gunzip &> /dev/null
then
  echo "gunzip not found, downloading"
  apt-get install gunzip
fi
# wget not found
if ! command -v wget &> /dev/null
then
  echo "wget not found, downloading"
  apt-get install wget
fi
# git not found
if ! command -v git &> /dev/null
then
  echo "git not found, downloading"
  apt-get install git
fi

### utilities end ###

# bin init
if ! [[ -d "./bin" ]]
then
  mkdir "$SCRIPT_DIR/bin"
fi

# Install kics
if ! [[ -f "$SCRIPT_DIR/bin/kics" ]]
then
  wget https://github.com/Checkmarx/kics/releases/download/v1.4.7/kics_1.4.7_linux_x64.tar.gz -O "$SCRIPT_DIR/bin/kics.tar.gz"
  tar -xf "$SCRIPT_DIR/bin/kics.tar.gz" -C "$SCRIPT_DIR/bin/kics/"
  rm "$SCRIPT_DIR/bin/kics.tar.gz"
fi

# Install keyfinder
pip3 install -r ./requirements.txt -q
if ! [[ -d "$SCRIPT_DIR/bin/keyfinder" ]]
then
  git clone https://github.com/CERTCC/keyfinder.git "$SCRIPT_DIR/bin/keyfinder/"
fi

# Install gitsecret
if ! [[ -d "$SCRIPT_DIR/bin/gitsecret" ]]
then
  git clone https://github.com/awslabs/git-secrets.git $SCRIPT_DIR/bin/gitsecret/
fi
sleep 1
cd "$SCRIPT_DIR/bin/gitsecret/" && make install

# Build our app
go build $SCRIPT_DIR