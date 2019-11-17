#!/bin/bash

set +e

echo "install liblua5.1"
sudo apt-get update
sudo apt-get install liblua5.1

echo "download and unzip sdl2 (2.0.9)"
curl -o sdl2.tar.gz http://libsdl.org/release/SDL2-2.0.9.tar.gz
tar -xvzf sdl2.tar.gz 1> /dev/null

echo "install sdl2 (2.0.9)"
cd SDL2-2.0.9
mkdir build
cd build
../configure
make
sudo make install
cd ../..

echo "remove sdl2 dir"
rm -rf SDL2-2.0.9

echo "install dep"
go get -u github.com/golang/dep/cmd/dep