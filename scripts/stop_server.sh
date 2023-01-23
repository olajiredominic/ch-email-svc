#!/bin/bash
kill -9 $(lsof -t -i:5004) > /dev/null 2> /dev/null < /dev/null &
