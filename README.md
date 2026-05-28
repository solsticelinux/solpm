# SOLPM
## What is Solpm?
It's a **simple** Package manager for **SolsticeOS**, Currently It will be written in bash but we will migrate it to another language if possible.

## Why?
Sadly, I only know how to code in **Bash** So I decided to use my current Skills to make a rough prototype of how Solpm could be.

## Binary or source ?
As of now I'll be working on Solpm's abillity to work with source code, and binaries are somewhat of an afterthought unless it is requested by the community.

## What is the recipe format?
The recipe format is also gonna be written in **Bash Script**, So each recipe is gonna be something like ``linux.sh``
``` bash
PKG="linux"
VER="7.1-rc5"
DEP=""
SRC="https://cdn.kernel.org/pub/linux/kernel/v7.x/linux-${VER}.tar.xz"

pull() {
    wget $SRC
    tar -xf linux-${VER}.tar.xz
    cd linux-${VER}
}

build() {
    make defconfig
    make -j$(nproc)
    make modules_install
    make install
}

pull
build
```
### **This is only an example of how our recipes could look like.**
