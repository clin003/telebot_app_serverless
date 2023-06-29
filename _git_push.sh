#!/bin/bash

VERSION=0.1.12
APPNAME=telebot_app_serverless
echo "package baicai" > ./baicai/version.go
echo "const(APP_NAME = \"${APPNAME}\"" >> ./baicai/version.go
echo "APP_VERSION = \"${VERSION}\")" >> ./baicai/version.go
go fmt ./baicai


#git init #
git add .
git commit -m "v${VERSION} debug"
#git remote add gitee git@gitee.com:lyhuilin/${APPNAME}.git #
#git remote add github git@github.com:clin003/${APPNAME}.git #
#git branch -M main #
git push -u gitee main
git push -u github main

git tag "v${VERSION}"
git push --tags  -u github main
git push --tags  -u gitee main
# git remote -v