#!/bin/bash
echo "0m test start"

sleep 10                    # 10s group a, b generate alert instance
echo "10s groupA instance0, groupB instance0 down"
docker stop group-a-node-0
docker stop group-b-node-0

sleep 5                     # 15s group a alert instance added
echo "15s groupA instance1 down"
docker stop group-a-node-1 

sleep 45                    # 40s groupA, groupB alert notification by group_wait
echo "30s groupA, groupB alert notification by group_wait"


sleep 15                     # 105s group a alert instance added
echo "1m 15s instance1 down"
docker stop group-a-node-2

sleep 45
echo "2m groupA alert notification by group_interval"

sleep 120
echo "4m groupB alert notification by repeat_interval"

sleep 60
echo "5m groupA alert notification by repeat_interval"


sleep 120
echo "7m test finished"
docker compose up group-a-node-0 -d
docker compose up group-a-node-1 -d
docker compose up group-a-node-2 -d
docker compose up group-b-node-0 -d