#!/bin/bash

# 编译 Go 程序
echo "Building the Go server..."
go build -o server/server main.go
if [ $? -ne 0 ]; then
  echo "Build failed! Exiting..."
  exit 1
fi

# 查找占用 8888 端口的进程
echo "Killing the process using port 8888..."
PID=$(lsof -t -i:8888)
if [ -n "$PID" ]; then
  kill -9 $PID
  echo "Killed process $PID on port 8888"
else
  echo "No process found on port 8888"
fi

# 启动服务器并将输出重定向到日志文件
echo "Starting the server..."
./server/server >> ./server/server.log 2>&1 &
if [ $? -eq 0 ]; then
  echo "Server started successfully!"
else
  echo "Failed to start the server!"
  exit 1
fi
