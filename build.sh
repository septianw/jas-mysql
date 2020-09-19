#!/bin/bash

APIVERSION=$(cat version.go | grep APIVersion | cut -d" " -f4 | sed "s/\"//g")
VERSION=$(cat VERSION);
COMMIT=$(git rev-parse --short HEAD);

WRITTENVERSION=$APIVERSION'-'$VERSION'-'$COMMIT

sed -i "s/versionplaceholder/"$WRITTENVERSION"/g" version.go

mkdir bungkus
go build -buildmode=plugin -ldflags="-s -w" -o bungkus/database.so
cp -Rvf LICENSE CHANGELOG bungkus
mv bungkus database
tar zcvvf database-$WRITTENVERSION.tar.gz database
rm -Rvf database
