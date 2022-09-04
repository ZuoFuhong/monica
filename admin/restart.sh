#!/bin/sh

PROJ_PATH=`pwd`
SERVER_NAME=`basename $PROJ_PATH`

PIDS=`ps -ef | grep -v grep | grep "$SERVER_NAME" |awk '{print $2}'`
if [ -n "$PIDS" ]; then
    echo "ERROR: The $SERVER_NAME already started!"
    echo "PID: $PIDS"
    #exit 1
    kill -9 $PIDS
fi

./$SERVER_NAME --conf conf/app_prod.yaml > info.log 2>&1 &
echo "done"