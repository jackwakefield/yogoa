#!/usr/bin/env bash
YOGA_VERSION=1.8.0

cd "$(dirname "$0")"

wget -q https://github.com/facebook/yoga/archive/$YOGA_VERSION.tar.gz -O $YOGA_VERSION.tar.gz
tar xzvf $YOGA_VERSION.tar.gz

rm -rf ../yoga
mkdir -p ../yoga
mv yoga-$YOGA_VERSION/yoga/*.{cpp,h} ../yoga/
mv yoga-$YOGA_VERSION/LICENSE ../yoga/

cd ../yoga
patch < ../scripts/yoga.patch
cd ../scripts

rm -rf yoga-$YOGA_VERSION
rm -f $YOGA_VERSION.tar.gz