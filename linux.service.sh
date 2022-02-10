#!/bin/sh
echo Starting github-backup service. 
date -I

echo --------------------------------------------------------------
echo "   Updating the service"
echo --------------------------------------------------------------
git pull

echo --------------------------------------------------------------
echo "   Updating npm packages"
echo --------------------------------------------------------------
npm i

echo --------------------------------------------------------------
echo "   Compiling typescript code" 
echo --------------------------------------------------------------
npm run compile

echo --------------------------------------------------------------
echo "   Starting backup service"
echo --------------------------------------------------------------
npm run start 

echo service finished.













