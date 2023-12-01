#!/bin/bash
rm -rf /tmp/kashd-temp

kashd --simnet --appdir=/tmp/kashd-temp --profile=6061 &
KASHD_PID=$!

sleep 1

orphans --simnet -alocalhost:16511 -n20 --profile=7000
TEST_EXIT_CODE=$?

kill $KASHD_PID

wait $KASHD_PID
KASHD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Kashd exit code: $KASHD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $KASHD_EXIT_CODE -eq 0 ]; then
  echo "orphans test: PASSED"
  exit 0
fi
echo "orphans test: FAILED"
exit 1
