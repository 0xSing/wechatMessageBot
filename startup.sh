#!/bin/bash
SERVER_NAME="wechatBot"

make wechatBot

screen -dmS syn_$CHAIN_ID bash -c "./build/bin/$SERVER_NAME; exec bash"

