#!/bin/bash
set -e

make
go install
mkdir -p ~/.hcc/clarinet/
cp ./complete ~/.hcc/clarinet/

if grep -Fxq  "source ~/.hcc/clarinet/complete" ~/.bashrc
then
	source ./complete
else
	echo "source ~/.hcc/clarinet/complete" >> ~/.bashrc
	source ./complete
fi
