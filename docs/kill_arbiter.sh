#!/bin/bash
PROCESS_NAME="arbiter"
if pgrep -x "$PROCESS_NAME" > /dev/null; then
    echo -e "\033[1;34mINFO:\033[0m Stopping $PROCESS_NAME..."
    pkill -x "$PROCESS_NAME"
    if pgrep -x "$PROCESS_NAME" > /dev/null; then
        echo -e "\033[1;31mERROR:\033[0m Failed to stop $PROCESS_NAME."
    else
        echo -e "\033[1;32mINFO:\033[0m $PROCESS_NAME has been stopped successfully."
    fi
else
    echo -e "\033[1;33mWARNING:\033[0m $PROCESS_NAME is not running."
fi