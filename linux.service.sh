$CURRENT_DATE = `date`
echo Starting github-backup service. $CURRENT_DATE

npm i
npm run compile
npm run start 

echo service finished.













