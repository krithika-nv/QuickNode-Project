#!/bin/bash

# create working directory
DIR=/root/daemon
if [ ! -d "$DIR" ]; then
mkdir -p /root/daemon
echo "Created working directory."
fi

# create systemd service
FILE=/etc/systemd/system/daemon.service
if [ ! -f "$FILE" ]
then
    mv /tmp/daemon.service /etc/systemd/system/
    echo "Created systemd service."
    systemctl daemon-reload
    systemctl enable daemon.service
fi
