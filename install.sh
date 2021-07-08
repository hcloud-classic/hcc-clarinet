#!/bin/bash
set -e

mkdir -p ~/.hcc/clarinet/
cp ./complete ~/.hcc/clarinet/

if grep -Fxq  ". ~/.hcc/clarinet/complete" ~/.bashrc
then
	. ~/.hcc/clarinet/complete
else
	echo ". ~/.hcc/clarinet/complete" >> ~/.bashrc
	. ~/.hcc/clarinet/complete
fi

echo "Install Finished"
echo "run command below to use clarinet auto completion"
echo "    source ~/.bashrc"
