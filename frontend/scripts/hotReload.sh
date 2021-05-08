#!/bin/bash
yarn build || echo "Build failed, entering hot reload loop. Now fix the problem dummy"
while true; do

inotifywait -e move,modify,create,delete -r ./ --exclude node_modules/ && \
    yarn buildDev

done
