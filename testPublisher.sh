#!/bin/bash

cd Publisher
go build publisher.go
cd ..
./Publisher/publisher -j ./Publisher/json -ms 300
