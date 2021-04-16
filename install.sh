#!/bin/bash
set -e

make
go install
mkdir -p ~/.hcc/clarinet/
cp ./complete ~/.hcc/clarinet/

if grep -Fxq  "source ~/.hcc/clarinet/complete" ~/.bashrc
then
	source ~/.hcc/clarinet/complete
else
	echo "source ~/.hcc/clarinet/complete" >> ~/.bashrc
	source ~/.hcc/clarinet/complete
fi

echo "Install Finished"
echo "run command below to use clarinet auto completion"
echo "    source ~/.bashrc"
