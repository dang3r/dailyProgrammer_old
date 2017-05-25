#!/usr/bin/env bash
#
# Posts on reddit require you to prepend 4 spaces to prevent spoilers

FILE=$1
sed -e 's/^/    /' "$FILE" > "$FILE".reddit
