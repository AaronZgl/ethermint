#!/bin/bash

tmux kill-session -t resource_recorder 1> /dev/null; tmux new -s resource_recorder -d; tmux send-keys -t resource_recorder "while true; do (date '+%s' && ps -e -o pcpu,pmem,args --sort=pcpu | grep \"ethermint\" | grep -v grep | cut -d\" \" -f1-5 | tail) >> resource.log; sleep 1; done" C-m

cd .. && go get && go build && cd emint
bash init.sh 2> /dev/null

sleep 10

../benchmarking s -c $1

echo "cool-down period: 30 s"
sleep 30

tmux kill-session -t resource_recorder

START=$(cat start.txt)
END=$(cat end.txt)
../benchmarking a -s $START -e $END

rm start.txt end.txt