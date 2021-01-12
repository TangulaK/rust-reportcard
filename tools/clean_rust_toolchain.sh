#!/bin/bash
cd ~/.rustup/toolchains/
for i in nightly-*; do rustup toolchain remove ${i/%-x86_64-unknown-linux-gnu}; done
