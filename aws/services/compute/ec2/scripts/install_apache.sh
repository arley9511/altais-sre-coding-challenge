#!/bin/bash
# ----------------------------------------------------------------------------
# user-data.bash
# ----------------------------------------------------------------------------
export LANG=en_US.UTF-8
export LC_ALL=${LANG}

yum -y update

yum install -y httpd.x86_64

systemctl start httpd.service

systemctl enable httpd.service

echo "Hello World from $(hostname -f)" > /var/www/html/index.html
