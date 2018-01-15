#!/bin/sh

# install dep
go get -u github.com/golang/dep/cmd/dep

# # download and unzip sdl2
# curl -o sdl2.tar.gz http://libsdl.org/release/SDL2-2.0.7.tar.gz
# tar -xvzf sdl2.tar.gz

# # install sdl2
# cd SDL2-2.0.7
# mkdir build
# cd build
# ../configure
# make
# sudo make install
# cd ../..

# # remove sdl2 dir
# rm -rf SDL2-2.0.7

sudo apt install libsdl2

# install lua
sudo apt-get install liblua5.1
