#!/bin/bash

sudo su root
cd /home/admin/streamingMedia/
git pull origin -m
chmod 777 /home/admin/streamingMedia/mediaServer/mediaServer

