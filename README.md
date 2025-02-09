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
Provider为提供下载数据的接口，目前支持的Provider有：Domp4Provider（网站地址：[https://www.ddmp4.cc](https://www.ddmp4.cc)）
如想扩展，实现`ParseURLs(URL string, CurrentEp int) ([]string, int, error)`接口

> **注意：provider请合理使用，来源均为公开互联网可使用的资源，禁止高频爬取造成网站压力，遵守网站包括不限于robots.txt的相关规定**

### Downloader
Downloader为下载数据的接口，目前支持的Downloader有：ThunderDownloader
如想扩展，实现`SendTask(task Task) error`接口

### TvTask
TvTask为追更任务，支持通过管理页面添加追剧任务，支持手动追剧，支持定时追剧

## 运行

### 运行准备
- 修改`config/config.yaml`内容，配置管理后台账号密码、数据库链接信息、迅雷地址
- 创建数据库、根据`sql/sql.sql`创建数据表

### 本地运行
- 执行`go run main.go`

### 定时追更配置
`pkg/cron/cron.go`中配置定时任务

### docker
- 目前已经提供了`Dockfile`,了解Docker的可以自行部署
- `docker-compose`方式部署暂未impl

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