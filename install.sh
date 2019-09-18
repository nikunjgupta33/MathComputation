#!/bin/bash
echo "Downloading And Installing Project Dependencies..."

set -e
export GOPATH="$GOPATH:$PWD"

#Using by Math Computation Service
go get "github.com/spf13/viper"
go get "github.com/iris-contrib/middleware/cors"
go get "github.com/kataras/iris"
go get "github.com/kataras/iris/context"

#Using by IRIS Internally
go get "github.com/CloudyKit/jet"
go get "github.com/Joker/jade"
go get "github.com/ryanuber/columnize"
go get "github.com/aymerick/raymond"
go get "github.com/flosch/pongo2"
go get "github.com/eknkc/amber"
go get "golang.org/x/crypto/acme/autocert"

if [ $? -ne 0 ]
then
    echo
    echo "Failed"
    exit 1
fi
    echo "Done"