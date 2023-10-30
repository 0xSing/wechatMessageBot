#!/bin/bash
SERVER_NAME="wechatBot"

while read -r pid; do
  if [[ -n "$pid" ]]; then
    kill -9 "$pid"
    echo "PID $pid is killed"
  fi
done <<< $(ps -ef | grep ./build/bin/$SERVER_NAME | grep -v grep | awk '{print $2}')

screen -wipe