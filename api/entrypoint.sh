#!/bin/sh

set -e

if [ -e /usr/bin/server ]
    then
        /usr/bin/server
    else
        exec "$@"
fi