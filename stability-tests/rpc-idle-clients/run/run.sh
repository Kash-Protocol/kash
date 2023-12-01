#!/bin/bash
rm -rf /tmp/kashd-temp

NUM_CLIENTS=128
kashd --devnet --appdir=/tmp/kashd-temp --profile=6061 --rpcmaxwebsockets=$NUM_CLIENTS &
KASHD_PID=$!
KASHD_KILLED=0
function killKashdIfNotKilled() {
  if [ $KASHD_KILLED -eq 0 ]; then
    kill $KASHD_PID
  fi
}
trap "killKashdIfNotKilled" EXIT

sleep 1

rpc-idle-clients --devnet --profile=7000 -n=$NUM_CLIENTS
TEST_EXIT_CODE=$?

kill $KASHD_PID

wait $KASHD_PID
KASHD_EXIT_CODE=$?
KASHD_KILLED=1

echo "Exit code: $TEST_EXIT_CODE"
echo "Kashd exit code: $KASHD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $KASHD_EXIT_CODE -eq 0 ]; then
  echo "rpc-idle-clients test: PASSED"
  exit 0
fi
echo "rpc-idle-clients test: FAILED"
exit 1
