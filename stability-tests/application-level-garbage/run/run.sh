#!/bin/bash
rm -rf /tmp/kashd-temp

kashd --devnet --appdir=/tmp/kashd-temp --profile=6061 --loglevel=debug &
KASHD_PID=$!
KASHD_KILLED=0
function killKashdIfNotKilled() {
    if [ $KASHD_KILLED -eq 0 ]; then
      kill $KASHD_PID
    fi
}
trap "killKashdIfNotKilled" EXIT

sleep 1

application-level-garbage --devnet -alocalhost:16611 -b blocks.dat --profile=7000
TEST_EXIT_CODE=$?

kill $KASHD_PID

wait $KASHD_PID
KASHD_KILLED=1
KASHD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Kashd exit code: $KASHD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $KASHD_EXIT_CODE -eq 0 ]; then
  echo "application-level-garbage test: PASSED"
  exit 0
fi
echo "application-level-garbage test: FAILED"
exit 1
