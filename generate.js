const child = require("child_process");
const fs = require("fs");

const output = child
    .execSync(`git log --format=%B%H----DELIMITER----`)
    .toString("utf-8");

const commitsArray = output
    .split("----DELIMITER----\n")
    .map(commit => {
        const [message, sha] = commit.split("\n");

        return { sha, message };
    })
    .filter(commit => Boolean(commit.sha));

const currentChangelog = fs.readFileSync("./CHANGELOG.md", "utf-8");
const currentVersion = String(require("./release_version.json").version);
const commitPath = String(require("./release_version.json").origin);

// 用version和时间作为release 标记
let newChangelog = `# Version ${currentVersion} (${
    new Date().toISOString().split("T")[0]
})\n\n`;

const features = [];
const bugfixes = [];

// 分别维护features和bugfixes的内容，并将message和commit的链接进行绑定
commitsArray.forEach(commit => {
    if (commit.message.startsWith("features: ")) {
        features.push(
            `* ${commit.message.replace("features: ", "")} ([${commit.sha.substring(
                0,
                6
            )}](${commitPath}/${
                commit.sha
            }))\n`
        );
    }
    if (commit.message.startsWith("bugfixes: ")) {
        bugfixes.push(
            `* ${commit.message.replace("bugfixes: ", "")} ([${commit.sha.substring(
                0,
                6
            )}](${commitPath}/${
                commit.sha
            }))\n`
        );
    }
});

if (features.length) {
    newChangelog += `## Features\n`;
    features.forEach(feature => {
        newChangelog += feature;
    });
    newChangelog += '\n';
}

if (bugfixes.length) {
    newChangelog += `## Bugfixes\n`;
    bugfixes.forEach(bugfix => {
        newChangelog += bugfix;
    });
    newChangelog += '\n';
}

// prepend the newChangelog to the current one
fs.writeFileSync("./CHANGELOG.md", `${newChangelog}${currentChangelog}`);

