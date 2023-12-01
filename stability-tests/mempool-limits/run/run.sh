#!/bin/bash

APPDIR=/tmp/kashd-temp
KASHD_RPC_PORT=29587

rm -rf "${APPDIR}"

kashd --simnet --appdir="${APPDIR}" --rpclisten=0.0.0.0:"${KASHD_RPC_PORT}" --profile=6061 &
KASHD_PID=$!

sleep 1

RUN_STABILITY_TESTS=true go test ../ -v -timeout 86400s -- --rpc-address=127.0.0.1:"${KASHD_RPC_PORT}" --profile=7000
TEST_EXIT_CODE=$?

kill $KASHD_PID

wait $KASHD_PID
KASHD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Kashd exit code: $KASHD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $KASHD_EXIT_CODE -eq 0 ]; then
  echo "mempool-limits test: PASSED"
  exit 0
fi
echo "mempool-limits test: FAILED"
exit 1
