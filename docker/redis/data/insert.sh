#!/bin/sh
#
FILE=/usr/local/initial_data.txt

insertData(){
    echo "Inserting data..."
    { cat ${FILE} || echo status="$?"; } | redis-cli -h redis --pipe
    echo "Terminou data..."   
}

insertData