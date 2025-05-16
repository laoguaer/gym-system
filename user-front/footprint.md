#配置代理
https://github.com/Elegycloud/clash-for-linux-backup

#安装docker
apt install docker.io
docker pull mysql:8.0
docker pull redis:7.0
docker run --name mysql -e MYSQL_ROOT_PASSWORD=123456 -p 3306:3306 -d mysql:8.0
docker run --name redis -p 6379:6379 -d redis:7.0

#安装Node.js
https://nodejs.org/en/download
nvm install 22
nvm use 22
npm install -g pnpm
