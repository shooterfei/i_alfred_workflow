export PATH=/usr/local/opt/libressl/bin:$PATH

export LDFLAGS="-L/usr/local/opt/libressl/lib"
export CPPFLAGS="-I/usr/local/opt/libressl/include"


sudo luarocks install luasocket
