# Github backup script

This script will backup your public github repositories to a local directory.

## ¿How to use it?

`Download the script` and place it in a `directory of your choice` and create the following directory:

    /path/to/script
    ├── github-backup-script
    └── Resources 
        ├── .env
        └── git.json

In the `.env` file, you must place your github username and `the path to the backup directory`.

The `token is optional`, just leave with "" if you don't want to use it.

```.env
github.backup.token=github_token
github.backup.path=/your/backup/directory
```

The `git.json` file is used tell the script which repositories you want to backup.

```json
[
    {
        "username": "github_username",
        "repositories": [
            "repository1",
            "repository2",
            "repository3"
        ]
    },
    {
        "username": "github_username2",
        "repositories": [
            "repository1"
        ]
    }
]
```


Make the script executable Then, run it:

```bash
chmod +x github-backup-script
./github-backup-script
```


# Github data fetch script

This script will fetch your public Github repository names 
and create the `git.json` file for you.


## ¿How to use it?

`Download the script` and place it in a `directory of your choice` and create the following directory:

    /path/to/script
    ├── github-data-fetch
    └── Resources 
        └── .env

In the `.env` file, you must place the `usernames of the accounts you want to backup`.

The `token is not necessary`, just leave with "".


```.env
github.fetch.usernames=akrck02,torvalds
github.fetch.token=-
github.fetch.json.path=/place/for/git.json/file
```

Make the script executable Then, run it:

```bash
chmod +x github-data-fetch
./github-data-fetch
```