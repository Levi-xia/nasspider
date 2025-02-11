# Nas追剧神器

## 简介
此项目为[Nas追剧神器](https://github.com/Levi-xia/nasspider)，提供图形化管理界面，旨在能够实现自动追更功能，能够实现根据任务配置，自动解析provider的下载数据，调用downloader进行下载

<img src="https://github.com/Levi-xia/nasspider/blob/main/img/one.jpg">

## 功能
- 新增追更任务
- 手动追更
- 定时追更

## 实现
<img src="https://github.com/Levi-xia/nasspider/blob/main/img/two.png">

### Provider
Provider为提供下载数据的接口，目前支持的Provider有：
- Domp4Provider（网站地址：[https://www.ddmp4.cc](https://www.ddmp4.cc)）

> 如想扩展，实现`ParseURLs(URL string, CurrentEp int) ([]string, int, error)`接口

> ⚠️provider请合理使用，来源均为公开互联网可使用的资源，禁止高频爬取造成网站压力，遵守网站包括不限于robots.txt的相关规定

### Downloader
Downloader为下载数据的接口，目前支持的Downloader有：
- ThunderDownloader

> 预期将会支持QB、Aria2等Downloader，如想扩展，实现`SendTask(task Task) error`接口

### TvTask
TvTask为追更任务，支持通过管理页面添加追剧任务，支持手动追剧，支持定时追剧

## 运行

### 本地运行
- 修改`config/config.yaml`内容，配置管理后台账号密码、数据库链接信息、迅雷地址
- 执行`go run main.go`
- `pkg/cron/cron.go`中配置定时任务

### docker
- 目前已经提供了`Dockfile`,了解Docker的可以自行部署
- 如果你是NAS发烧友，但是不太懂开发，那么本项目也提供`docker-compose`方式一键部署
```yaml
version: '3'

networks:
  nas-spider-network:
    driver: bridge

services:
  mysql:
    image: mysql:5.7
    container_name: nas-spider-mysql
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: '<数据库密码>'
      MYSQL_DATABASE: 'nas-spider'
      MYSQL_ROOT_HOST: '%'
    volumes:
      - ./data/mysql:/var/lib/mysql
    networks:
      nas-spider-network:
        aliases:
          - nas-spider-network-mysql
    command:
      - --character-set-server=utf8mb4
      - --collation-server=utf8mb4_unicode_ci
      - --default-authentication-plugin=mysql_native_password

  web:
    image: registry.cn-beijing.aliyuncs.com/levicy/nas-spider:latest
    container_name: nas-spider-web
    restart: always
    environment:
      MYSQL_HOST: 'nas-spider-network-mysql'
      MYSQL_PORT: '3306'
      MYSQL_USER: 'root'
      MYSQL_PASSWORD: '<数据库密码>'
      SERVER_PORT: '<服务器端口号>'
      THUNDER_HOST: 'http://<迅雷地址>'
      THUNDER_PORT: '<迅雷端口号>'
      ADMIN_USERNAME: '<后台账号>'
      ADMIN_PASSWORD: '<后台密码>'
    ports:
      - "<映射宿主机端口号>:<服务器端口号>"
    networks:
      nas-spider-network:
        aliases:
          - nas-spider-network-web
    depends_on:
      - mysql
```

## 使用
以目前支持的，provider为domp4，downloader为thunder举例：
- 点击新增，添加追剧任务
- 在弹出层中填写信息
- 点击保存，保存成功
- 点击手动追更，手动追更

<img src="https://github.com/Levi-xia/nasspider/blob/main/img/three.jpg" style="width: 50%">

## 其他
### 迅雷配置下载文件夹
TvTask中的`download_path`可以设置为默认的`/downloads/[目标文件夹/]`(目标文件夹可以自动新建),

如果想要修改（以飞牛Nas为例）：
- 应用管理关闭迅雷，docker中关闭迅雷容器，修改文件映射
- 设置迅雷访问权限
- 重新运行容器
- 应用管理运行迅雷


## 免责声明:
1. 该项目设计和开发仅供学习、研究和安全测试目的。请于下载后 24 小时内删除, 不得用作任何商业用途, 文字、数据及图片均有所属版权, 如转载须注明来源。
2. 使用本程序必循遵守部署服务器所在地区的法律、所在国家和用户所在国家的法律法规。对任何人或团体使用该项目时产生的任何后果由使用者承担。
3. 作者不对使用该项目可能引起的任何直接或间接损害负责。作者保留随时更新免责声明的权利，且不另行通知。