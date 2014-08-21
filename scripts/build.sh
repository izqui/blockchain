#!/bin/bash
docker build -t izqui/blockchain .
boot2docker down
boot2docker up