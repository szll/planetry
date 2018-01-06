#!/bin/sh

# install dep
go get -u github.com/golang/dep/cmd/dep

# download and unzip sdl2
curl -o sdl2.tar.gz http://libsdl.org/release/SDL2-2.0.7.tar.gz
tar -xvzf sdl2.tar.gz

# install sdl2
cd SDL2-2.0.7
mkdir build
cd build
../configure
make
sudo make install
cd ../..

# remove sdl2 dir
rm -rf SDL2-2.0.7

# download and unzip lua5.1
curl -o lua5.1.tar.gz https://www.lua.org/ftp/lua-5.1.5.tar.gz
tar -xvzf lua5.1.tar.gz

# install lua5.1
cd lua-5.1.5
make
make linux
sudo make install

sudo ln /usr/local/lib/liblua.a /usr/local/lib/liblua5.1.a
sudo ln /usr/local/include/lua.h /usr/local/include/lua5.1.h 
sudo ln /usr/local/include/luaconf.h /usr/local/include/luaconf5.1.h
sudo ln /usr/local/include/lualib.h /usr/local/include/lualib5.1.h
sudo ln /usr/local/include/lauxlib.h /usr/local/include/lauxlib5.1.h
sudo ln /usr/local/include/lua.hpp /usr/local/include/lua5.1.hpp 
cd ..

# remove lua5.1 dir
rm -rf lua-5.1.5