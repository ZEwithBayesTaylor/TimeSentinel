wrk -t50 -c200 -d30s --script=xtimer_create_timer.lua --latency "http://127.0.0.1:9001/xtimer/createTimer"
