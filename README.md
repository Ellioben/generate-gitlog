
# node to generate changelog.md
base on git commit message to generate changelog.md
## config process 🍿
1. brew install node
2. cd generate-gitlog
3. create CHANGELOG.md `vim CHANGELOG.md`
4. config release_version.json

## generate changelog & update chengelog
git message format <git commit -m "bugfixes: xxxxx"> or <git commit -m "features: xxxxx"> 

`node generate.js`
