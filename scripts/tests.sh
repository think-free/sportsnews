#!/bin/bash

# Run the test for the main app
cd sportsnews
go test ./...
cd ..

# For each providers in folder, run the test
for d in providers/* ; do
    cd $d
    go test ./...
    cd ../..
done
