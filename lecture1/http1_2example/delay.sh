#!/usr/bin/env bash

LATENCY="${1-0}"
ERROR="${2-0}"
REAL_LATENCY=$((LATENCY/2))

echo "Latency: $LATENCY msec, loss:$ERROR%"

sudo tc qdisc del dev lo root
sudo tc qdisc add dev lo root handle 1:0 netem delay "${REAL_LATENCY}msec" loss "${ERROR}%"
