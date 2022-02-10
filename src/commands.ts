const { execSync } = require("child_process");

export class Commands {

    public static MakeDirectory(directory: string): void {
        execSync(`mkdir -p ${directory}`,{stdio: 'inherit'});
    }

    public static gitClone(directory: string, user: string, repo: string): void {
        execSync(`cd ${directory} && git clone https://github.com/${user}/${repo}`,{stdio: 'inherit'});
    }

    public static gitPull(directory: string): void {
        execSync(`cd ${directory} && git pull`,{stdio: 'inherit'});
    }

    

}