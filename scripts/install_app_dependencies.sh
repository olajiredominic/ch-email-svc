#!/bin/bash
echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.profile
source ~/.profile

cd /home/ubuntu
chmod 777 -R ./email_svc
cd ./email_svc
go build -o email_svc ./cmd
