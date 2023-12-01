#!/bin/bash
rm -rf /tmp/kashd-temp

kashd --devnet --appdir=/tmp/kashd-temp --profile=6061 --loglevel=debug &
KASHD_PID=$!

sleep 1

rpc-stability --devnet -p commands.json --profile=7000
TEST_EXIT_CODE=$?

kill $KASHD_PID

wait $KASHD_PID
KASHD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Kashd exit code: $KASHD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $KASHD_EXIT_CODE -eq 0 ]; then
  echo "rpc-stability test: PASSED"
  exit 0
fi
echo "rpc-stability test: FAILED"
exit 1
