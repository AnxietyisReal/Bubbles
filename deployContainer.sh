#!/bin/bash
# Remotes into production server and deploys the container.

echo "Deploying container to production server..."
ssh ovhdocker "cd Bubbles && docker compose pull && docker compose up -d && exit"
echo "Done."

###EOF###