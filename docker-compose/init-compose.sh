#!/bin/bash

# use new version of tctl so that we can skip the prompt
tctl config set version next

for run in {1..60}; do
  sleep 1
  echo "now trying to register iWF system search attributes..."
  if tctl search-attribute  create -name IwfWorkflowType -type Keyword -y; then
    break
  fi
done
sleep 0.1
tctl search-attribute  create -name ObjectId -type Keyword -y
sleep 0.1
tctl search-attribute  create -name ObjectExecutionId -type Keyword -y
sleep 0.1
tctl search-attribute  create -name ObjectType -type Keyword -y
sleep 0.1
tctl search-attribute  create -name IwfGlobalWorkflowVersion -type Int -y
sleep 0.1
tctl search-attribute  create -name IwfExecutingStateIds -type Keyword -y
sleep 0.1
tctl namespace register default

tail -f /dev/null