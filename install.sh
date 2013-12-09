#!/bin/bash
# Sorry no GNU autotools, as simple as:
GOARM=5 go build
strip fan3xxnsa
mv -v fan3xxnsa /usr/local/bin
