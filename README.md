## 系统提供查考价值，直接使用未必能达到您想要的预期
### store-chat 一个支持私信和广播消息的即时通讯系统
* store-chat是利用go-zero搭建的一个即时通讯系统
* websocket服务可进行分布式部署，支持水平扩容（该架构使用了docker集群模式，建立多个websocket容器服务实现分布式）
* 使用redis的消息发布订阅进行推送接收消息，每个socket服务都会进行订阅redis
* 消息发布通过业务层 rpc.socket 服务将所有消息都推送到redis
* rpc的通讯是注册到etcd上的，通过etcd的注册服务发现服务进行通讯，并非直连RPC
* 兼容主从、单机的mysql服务，只需要设置对应配置文件即可
* mysql使用了gorm框架

### 系统搭建必备前提
#### 环境
* go: ^1.21
* mysql
* redis
* etcd
#### 配置文件存储
1. 请将所有 *.yaml.example 复制一份并命名 .example
2. 请将 192.168.33.10 ip替换成你服务器的docker的宿主机IP
3. 例如：在linux下安装的docker，通过 ifconfig docker0 找到宿主机IP
4. yaml位置
    * 项目/api/etc
    * 项目/rpc/socket/etc
    * 项目/socket/etc 
5. 以下是公共yaml位置
    * 项目/tools/yamls/*.yaml

#### 构建docker容器服务
1. ETCD服务（必须先搭起etcd服务，否则rpc无法连接到etcd注册） 
    * 项目/docker/etcd，执行 build.sh 脚本搭建服务
2. 项目服务，项目/docker/dev
    * .env:搭建容器的配置文件，请修改成你服务对应的文件目录
    * 各 shell 脚本描述
      * 若脚本没权限执行，先赋予脚本执行权限;
        * 命令: chmod +x *.sh
        * 需要赋予sh的目录；
          * 项目/docker/build
          * 项目/docker/dev
      * docker-build.sh: 构建镜像并构建容器且build服务文件，最后服务启动容器
      * docker-rebuild-compose.sh: 重新构建容器且build服务文件，最后启动服务容器
      * docker-restart.sh: 重新build各服务的 go文件，最后重启服务容器
      * docker-stop.sh: 停止所有 服务容器
    * 注意：所有服务的程序，都是通过 go-store-chat 容器进行build出来的
3. 若需要搭建mysql主从模式调试，也有快捷方式构建，（资源有限，up直接通过docker搭建了mysql主从）
   * 项目/docker/mysql, 执行 build.sh 脚本

#### mysql文件
1. 项目/docker/mysql/sql/*.sql


### 目录描述
* api:   api接口服务
* rpc:   grpc服务
* socket:   websocket服务
* socket-client:   仿照客户端进行websocket握手进程挂载，挂了两个自动回复的用户socket
* test:  开发调试
* dbs:   初始化数据库连接
* docker:   搭建服务
* model: 模型存放
  * mysqls: gorm框架模型存放
* tools: 公共code|公共变量|工具|公共类型|公共配置

### 功能演示
* 前端页面由 https://github.com/haostart/websocket-chat.git 提供,且改造开发的
* 有需要可以参考参考
* 演示地址；