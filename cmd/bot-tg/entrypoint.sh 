#!/bin/bash -e

exec > >(tee -a /var/log/app/entry.log | logger -t bot-tg -s 2>/dev/console) 2>&1

APP_ENV=${APP_ENV:-local}

echo "[$(date)] Running entrypoint script in the '${APP_ENV}' environment..."

CONFIG_FILE=./config/${APP_ENV}.yml

echo "[$(date)] Starting bot-tg..."
./bot-tg -config ${CONFIG_FILE} >>/var/log/app/bot-tg.log 2>&1
