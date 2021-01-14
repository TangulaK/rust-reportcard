#!/bin/bash
cd ~/.rustup/toolchains/
for i in *
do
    if [[ ${i} != "stable-x86_64-unknown-linux-gnu" ]]
    then
      rustup toolchain remove $i
    fi
done
