#!/bin/sh

# Start the background service
/usr/local/bin/poll &

# Optionally, wait a bit if your service needs to be up before the main app
sleep 2

# Run the passed command (from the final CMD or docker run)
exec "$@"