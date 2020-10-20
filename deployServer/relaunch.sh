#!/bin/bash

sudo su root
cd /home/admin/streamingMedia/
git pull origin main:main
chmod 777 /home/admin/streamingMedia/mediaServer/mediaServer

