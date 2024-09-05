BIN=sql-sample

go build -o $BIN main.go

./$BIN &

PID=$(ps aux | grep $BIN | grep -v grep | awk '{print $2}')

./grf-test value $PID -k Passwd

kill -9 $PID


