#!/bin/bash
rm -rf /tmp/kashd-temp

kashd --devnet --appdir=/tmp/kashd-temp --profile=6061 &
KASHD_PID=$!

sleep 1

infra-level-garbage --devnet -alocalhost:16611 -m messages.dat --profile=7000
TEST_EXIT_CODE=$?

kill $KASHD_PID

wait $KASHD_PID
KASHD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Kashd exit code: $KASHD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $KASHD_EXIT_CODE -eq 0 ]; then
  echo "infra-level-garbage test: PASSED"
  exit 0
fi
echo "infra-level-garbage test: FAILED"
exit 1
