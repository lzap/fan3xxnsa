#!/bin/bash
# Sorry no GNU autotools, as simple as:
go build
strip fan3xxnsa
mv fan3xxnsa ~/bin
