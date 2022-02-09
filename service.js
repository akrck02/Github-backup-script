const readerScript = require("./src/reader");
const {
    execSync
} = require("child_process");

const fs = require('fs');

function startApp() {
    try {

        console.log("############################################################################################");
        console.log("                           Github-backup-service by akrck02                                 ");
        console.log("############################################################################################\n");

        const params = readerScript.ReadParamsFile();
        checkParams(params);

        execSync(`mkdir -p ${params.directory}`);
        const repos = readerScript.ReadReposFile();
        const interval = +params.interval * 60 * 60 * 1000;
        console.log("interval: ", interval);

        cloneGithubRepos(params.directory, repos);

        setInterval(() => {
            console.log("Start backup at: ", new Date());
            cloneGithubRepos(params.directory, repos);
        }, interval);

    } catch (e) {
        console.log(e);
        console.log("[INFO] The script has failed. Please check the logs.\n");
        process.exit(1);
    }
}


function cloneGithubRepos(directory, reposFile) {

    let cloned = 0;
    let updated = 0;
    let failed = 0;

    for (const user in reposFile) {

        const repos = reposFile[user];
        const url = directory + "/" + user;
        execSync(`mkdir -p ${url}`);

        for (let repo of repos) {

            // if the directory exists, enter it and pull the repo
            if (fs.existsSync(url + "/" + repo)) {
                console.log("\n--------------------------------------------------------------------------------");
                console.log(`     Pulling ${user}/${repo}                              `);
                console.log("--------------------------------------------------------------------------------");

                execSync(`cd ${url}/${repo} && git pull`, {stdio: 'inherit'});
                updated++;
            } else {

                console.log("\n--------------------------------------------------------------------------------");
                console.log(`    Cloning ${user}/${repo}                              `);
                console.log("--------------------------------------------------------------------------------");

                //executing git clone on backup directory
                try {
                    execSync(`cd ${url} && git clone https://github.com/${user}/${repo}`, {
                        stdio: 'inherit'
                    });
                    cloned++;
                    console.log("[SUCCESS] ", repo, " cloned successfully. \n");
                } catch (error) {
                    console.log("[ERROR] Error cloning ----> ", repo);
                    failed++;
                }

            }
        }
    }

    console.log("\n--------------------------------------------------------------------------------");
    console.log(`    Cloned: ${cloned}                                                 `);
    console.log(`    Updated: ${updated}                                                 `);
    console.log(`    Failed: ${failed}                                                 `);
    console.log("--------------------------------------------------------------------------------");

}

/**
 * Check if the params are correct
 * @param {*} params 
 */
function checkParams(params) {

    if (params.token === undefined || params.token === "") {
        console.log("[WARNING] No token defined in params.json, only public repos will be backed up");
    }

    if (params.directory === undefined || params.directory === "") {
        console.log("[ERROR] No directory defined in params.json default directory is ~/backup");
        params.directory = "~/backup";
    }

    if(params.interval === undefined || params.interval === ""){
        console.log("[WARNING] No interval defined in params.json, defaulting to 30 minutes");
        params.interval = 30;
    }
}



startApp();