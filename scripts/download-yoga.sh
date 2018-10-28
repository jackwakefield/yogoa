#!/usr/bin/env bash
set -e

YOGA_VERSION=master

cd "$(dirname "$0")"

wget -q https://github.com/jackwakefield/yoga/archive/$YOGA_VERSION.tar.gz -O $YOGA_VERSION.tar.gz
tar xzvf $YOGA_VERSION.tar.gz

rm -rf ../pkg/yoga
mkdir -p ../pkg/yoga
mv yoga-$YOGA_VERSION/yoga/*.{cpp,h} ../pkg/yoga/
cp yoga-$YOGA_VERSION/LICENSE ../pkg/yoga/

rm -rf ../test/fixtures
mkdir -p ../test/fixtures
mv yoga-$YOGA_VERSION/gentest/fixtures ../test/
cp yoga-$YOGA_VERSION/LICENSE ../test/

rm -rf yoga-$YOGA_VERSION
rm -f $YOGA_VERSION.tar.gz
