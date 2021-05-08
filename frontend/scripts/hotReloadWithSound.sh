#!/bin/bash
cd "${0%/*}/.." # make sure cwd is correct for file paths
try_build () {
				yarn buildDev && succeed_sound || fail_sound
}

succeed_sound () {
				paplay ./scripts/hotReloadSounds/succeed.ogg
}

fail_sound () {
				paplay ./scripts/hotReloadSounds/failed.ogg --volume=65535
}

try_build
while true; do

inotifywait -e move,modify,create,delete -r ./ --exclude node_modules/ && \
    try_build

done
