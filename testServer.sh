#!/bin/bash

go build main.go
gnome-terminal -x ./main
sleep 1
gnome-terminal -x sh -c {
  cd Publisher &&
  go build publisher.go &&
  cd .. &&
  ./Publisher/publisher -j ./Publisher/json -ms 300;
}

