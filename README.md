
# node to generate changelog.md
base on git commit message to generate changelog.md
## config process 🍿
1. brew install node
2. cd generate-gitlog
3. vim CHANGELOG.md
4. release_version.json

## generate changelog & update chengelog
git message format <git message "bugfixes: xxxxx"> or <git message "features: xxxxx"> 

`node generate.js`
