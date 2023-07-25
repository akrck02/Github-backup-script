# Show environment variables
echo "Environment variables"
echo "---------------------------------------------------------"
echo " "
echo "GITHUB_REFRESH_INTERVAL: " $GITHUB_REFRESH_INTERVAL
echo "GITHUB_BACKUP_USER: "      $GITHUB_BACKUP_USER
echo "GITHUB_BACKUP_TOKEN: "     $GITHUB_BACKUP_TOKEN

echo "GITHUB_FETCH_USERNAMES: "  $GITHUB_FETCH_USERNAMES
echo "GITHUB_FETCH_TOKEN: "      $GITHUB_FETCH_TOKEN
echo "GITHUB_FETCH_JSON_PATH: "  $GITHUB_FETCH_JSON_PATH
echo " "
echo "---------------------------------------------------------"

# Remove old code
rm -rf Github-backup-script && echo "Removed old code"

# Download code from Github
git clone https://github.com/akrck02/Github-backup-script.git && echo "Downloaded code from Github"

# build go app
cd /app/Github-backup-script/backup && go build -o ../bin/backup/backup && echo "Built go backup script"
cd /app/Github-backup-script/data && go build -o ../bin/data/data && echo "Built go data script"

# remove source files
cd /app/Github-backup-script && rm -rf backup data .gitgnore readme.md .github && echo "Removed source files"

# add execution permissions
chmod +x /app/Github-backup-script/bin/backup/backup && echo "Added execution permissions for backup script"
chmod +x /app/Github-backup-script/bin/data/data && echo "Added execution permissions for data script"

# create .env files
mkdir -p /app/Github-backup-script/bin/data/Resources
mkdir -p /app/Github-backup-script/bin/backup/Resources
touch /app/Github-backup-script/bin/data/Resources/.env
touch /app/Github-backup-script/bin/backup/Resources/.env

# add environment variables
echo "github.fetch.usernames=$GITHUB_FETCH_USERNAMES"  >> /app/Github-backup-script/bin/data/Resources/.env
echo "github.fetch.token=$GITHUB_FETCH_TOKEN"          >> /app/Github-backup-script/bin/data/Resources/.env
echo "github.fetch.json.path=$GITHUB_FETCH_JSON_PATH"  >> /app/Github-backup-script/bin/data/Resources/.env

echo "github.backup.token=$GITHUB_BACKUP_TOKEN"        >> /app/Github-backup-script/bin/backup/Resources/.env
echo "github.backup.path=$GITHUB_BACKUP_PATH"          >> /app/Github-backup-script/bin/backup/Resources/.env

echo "FETCH .env file"
cat /app/Github-backup-script/bin/data/Resources/.env

echo " "
echo "BACKUP .env file"
cat /app/Github-backup-script/bin/backup/Resources/.env

echo " "

# run go app
cd /app/Github-backup-script/bin/data/ &&./data && cd /app/Github-backup-script/bin/backup/ && ./backup && echo "Github backup service is running!"