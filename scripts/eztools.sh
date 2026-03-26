#!/usr/bin/env bash
BIN="bin"
TMP="$BIN/tmp"
URL="https://download.ericzimmermanstools.com/net9"

echo "Download tools"

download () {
    if [ ! -f "$BIN/$1.dll" ] ; then
        echo "--- $1"
        wget "$URL/$1.zip"
		unzip -o "$1.zip" -d "$TMP"
    else
        return 1
    fi
}

install () {
    cp "$TMP/$1.dll" $BIN
    cp "$TMP/$1.runtimeconfig.json" $BIN
}

installMap () {
    cp -r "$TMP/$2/Maps" $BIN
    install "$2/$1"
}

mkdir -p $BIN $TMP

download "EvtxECmd" && installMap "EvtxECmd" "EvtxeCmd"
download "JLECmd" && install "JLECmd"
download "SBECmd" && install "SBECmd"

rm -rf $TMP

EZTOOLS=$(realpath $BIN)
export EZTOOLS
