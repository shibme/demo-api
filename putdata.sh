#!/bin/sh

curl -X PUT -H "Content-Type: application/json" -H "Authorization: Bearer demo_token" -d '{"data":"hello there!!"}' http://localhost:8888/data