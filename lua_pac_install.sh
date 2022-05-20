#!/usr/bin/env bash 

# export PATH=/usr/local/opt/libressl/bin:$PATH
# export LDFLAGS="-L/usr/local/opt/libressl/lib"
# export CPPFLAGS="-I/usr/local/Cellar/libressl/3.4.3/include"

packages=( luasocket lua-cjson2 "luasec OPENSSL_DIR=/usr/local/opt/libressl")

for pac in "${packages[@]}"
do
  luarocks install --lua-version 5.1 --local $pac
done

