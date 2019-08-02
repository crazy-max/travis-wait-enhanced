#!/bin/bash
set -e

export TRAVIS_WAIT_ENHANCED_VERSION=${TRAVIS_WAIT_ENHANCED_VERSION:-0.1.1}

echo "Installing travis-wait-enhanced..."
curl -LfsS -o /tmp/travis-wait-enhanced.zip "https://github.com/crazy-max/travis-wait-enhanced/releases/download/v${TRAVIS_WAIT_ENHANCED_VERSION}/travis-wait-enhanced_${TRAVIS_WAIT_ENHANCED_VERSION}_windows_x86_64.zip"
7z x /tmp/travis-wait-enhanced.zip -y -o/usr/bin/ travis-wait-enhanced.exe -r
travis-wait-enhanced --version
