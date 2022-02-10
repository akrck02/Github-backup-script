import { Commands } from "./commands";
import { Parameters } from "./configuration";
import { Configuration } from "./configuration";

const fs = require('fs');

export class Service {

    private cloned : number;
    private updated : number;
    private failed : number;

    private params : Parameters;
    private repos : any;

    /**
     * Create a new instance of the service
     * checking the compulsory parameters
     */
    constructor() {

        this.cloned = 0;
        this.updated = 0;
        this.failed = 0;

        try {
            this.params = Configuration.ReadParamsFile();
            this.checkParams();

            Commands.MakeDirectory(this.params.directory || "~/backup");
            this.repos = Configuration.ReadReposFile();
        } catch (e) {
            console.log(e);
            console.log("[INFO] The script has failed. Please check the logs.\n");
            process.exit(1);
        }

    }

    /**
     * Start the script execution
     */
    public start() {
        this.cloned = 0;
        this.updated = 0;
        this.failed = 0;

        this.showTitle();
        try {
            console.log("Start backup at: ", new Date());
            this.backup();
    
        } catch (e) {
            console.log(e);
            console.log("[INFO] The script has failed. Please check the logs.\n");
            process.exit(1);
        }
    
    }

    /**
     * Show the fancy title of the script
     */
    private showTitle() {
        console.log("############################################################################################");
        console.log("                           Github-backup-service by akrck02                                 ");
        console.log("############################################################################################\n");
    }

    /**
     * Check the compulsory parameters
     */
    private checkParams() {

        if (this.params.token === undefined || this.params.token === "") {
            console.log("[WARNING] No token defined in params.json, only public repos will be backed up");
        }
    
        if (this.params.directory === undefined || this.params.directory === "") {
            console.log("[ERROR] No directory defined in params.json default directory is ~/backup");
            this.params.directory = "~/backup";
        }
    
    }

    /**
     * Backup the repositories
     */
    private backup() {

        for (const user in this.repos) {
    
            const repos = this.repos[user];
            const url = this.params.directory + "/" + user;
            
            Commands.MakeDirectory(url);
            for (let repo of repos) {
    
                // if the directory exists, enter it and pull the repo
                if (fs.existsSync(url + "/" + repo)) {
                    this.updateRepository(url, user, repo);
                } else {
                    this.cloneRepository(url, user, repo);
                }
            }
        }
    
        console.log("\n--------------------------------------------------------------------------------");
        console.log(`    Cloned: ${this.cloned}                                                 `);
        console.log(`    Updated: ${this.updated}                                                 `);
        console.log(`    Failed: ${this.failed}                                                 `);
        console.log("--------------------------------------------------------------------------------");
    
    }

    /**
     * Clone a repository 
     * @param url The url of the repository
     * @param user The user of the repository
     * @param repo The repository name
     */
    private cloneRepository(url : string, user : string, repo : string) {
        console.log("\n--------------------------------------------------------------------------------");
        console.log(`    Cloning ${user}/${repo}                              `);
        console.log("--------------------------------------------------------------------------------");

        //executing git clone on backup directory
        try {
        
            Commands.gitClone(url, user, repo);
            this.cloned++;

            console.log("[SUCCESS] ", repo, " cloned successfully. \n");
        
        } catch (error) {
            this.failed++;
            console.log("[ERROR] Error cloning ----> ", repo);
        }

    }

    /**
     * Update a repository
     * @param url The url of the repository
     * @param user The user of the repository
     * @param repo The repository name
     */
    private updateRepository(url : string, user : string, repo : string) {
        console.log("\n--------------------------------------------------------------------------------");
        console.log(`     Pulling ${user}/${repo}                              `);
        console.log("--------------------------------------------------------------------------------");

        Commands.gitPull(`${url}/${repo}`);
        this.updated++;
    }
}


const service = new Service();
service.start();