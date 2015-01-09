echo "\n\n# Benchmarking Node.js"
node server.js &
PID=$!
sleep 2
ab -c 100 -n 10000 http://127.0.0.1:8000/
kill $PID
