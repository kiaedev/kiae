#!/bin/sh

FRONT_VERSION="v0.0.12"

curl -Lf https://github.com/kiaedev/kiae-front/releases/download/${FRONT_VERSION}/kiae-front-dist.tar.gz -o kiae-front-dist.tar.gz
tar zxvf kiae-front-dist.tar.gz && rm -rf kiae-front-dist.tar.gz