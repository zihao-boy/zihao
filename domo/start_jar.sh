#!/bin/bash

java -jar -Dspring.profiles.active=$2 target/service-$1.jar