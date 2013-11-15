#!/bin/bash
# Sorry no GNU autotools, as simple as:
go build
strip fan3xxnsa
mv -v fan3xxnsa /usr/local/bin
