#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

echo $DIR
aws cloudformation create-stack \
  --stack-name MyECSClusterStack \
  --template-body file://$DIR/go_app_template.yml \
  --parameters "$(cat $DIR/parameters.json)"
