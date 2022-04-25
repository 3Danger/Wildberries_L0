#!/bin/bash

fuser -k 4222/tcp &> /dev/null &
fuser -k 8080/tcp &> /dev/null &
