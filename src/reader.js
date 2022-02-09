const fs = require('fs');
const MEDIA_DIR = './media/';

function ReadParamsFile () {

    // Read the config file
    let params = {};
    let paramsFile;
    
    if (fs.existsSync(MEDIA_DIR + 'params.json')) {
        paramsFile = fs.readFileSync(MEDIA_DIR + 'params.json', 'utf8');
        params = JSON.parse(paramsFile);
    } else {
        throw '[ERROR] No params file found';
    }

    // Return the config
    return params;
}

function ReadReposFile () {

    // Read the repos file
    let repos = {};
    let reposFile;
    if (fs.existsSync(MEDIA_DIR + 'repos.json')) {
        reposFile = fs.readFileSync(MEDIA_DIR + 'repos.json', 'utf8');
        repos = JSON.parse(reposFile);
    } else {
        throw '[Github-backup][ERROR] No repos file found';
    }

    repos = JSON.parse(reposFile);

    // Return the repos
    return repos;

}

module.exports = {
    ReadParamsFile,
    ReadReposFile
};