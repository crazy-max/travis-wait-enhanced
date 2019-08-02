#!/bin/bash
set -e

export TRAVIS_WAIT_ENHANCED_VERSION=${TRAVIS_WAIT_ENHANCED_VERSION:-0.1.1}

echo "Installing travis-wait-enhanced..."
wget -qO- "https://github.com/crazy-max/travis-wait-enhanced/releases/download/v${TRAVIS_WAIT_ENHANCED_VERSION}/travis-wait-enhanced_${TRAVIS_WAIT_ENHANCED_VERSION}_linux_x86_64.tar.gz" | tar -zxvf - travis-wait-enhanced
mv travis-wait-enhanced /home/travis/bin/
travis-wait-enhanced --version
