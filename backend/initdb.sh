#!/bin/bash

touch sur.db
sqlite3 -init sur.db
.quit
sqlite3 sur.db < initdb.sql 
if [[ $? -eq 0 ]]; then 
    echo "Database created successfully"
else
    echo "Database creation failed"
fi
