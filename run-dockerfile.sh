#!/bin/bash
docker build -t coffee-orders .
docker run -dp 3000:3000 coffee-orders