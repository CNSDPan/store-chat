## Mysql主从库搭建
###### 描述
```text
搭建前，请参考 mysql/docker-compose.yml的容器构建配置:
command: mysql-master和mysql-slave都配置了相关配置
```
#### 主库(master)
###### 需要手动进入mysql-master容器创建一个账号给从库访问同步拉取数据的
```text
docker exec -ti mysql-master bash
# 登录mysql
mysql -uroot -proot

# 创建slave用户
create user 'sync'@'%' identified by '123456';
# 授予复制权限
grant replication slave,replication client on *.* to 'sync'@'%';
# 刷新权限
flush privileges;
# 记下File和Position的值,从库配置需求用到
SHOW MASTER STATUS;
+---------------+----------+--------------+------------------+-------------------+
| File          | Position | Binlog_Do_DB | Binlog_Ignore_DB | Executed_Gtid_Set |
+---------------+----------+--------------+------------------+-------------------+
| binlog.000007 |      767 |              |                  |                   |
+---------------+----------+--------------+------------------+-------------------+
1 row in set (0.03 sec)
```
```text
# 查看有多个是slave
SHOW SLAVE HOSTS;
```
#### 从库(slave)
###### 需要手动进入mysql-slave1容器创建一个账号给从库访问同步拉取数据的
```text
docker exec -ti mysql-slave1 bash
# 登录mysql
mysql -uroot -proot

# 先停止从数据库的复制线程
stop slave;

# 配置主从关系
change master to master_host='192.168.33.10',master_user='sync',master_password='123456',master_port=3306,master_log_file='binlog.000007', master_log_pos=767,master_connect_retry=30;

#开启从数据库的复制线程
start slave;

# 配置描述
master_host: mysql-master库IP
master_user: master库给slave使用的用户
master_password: master库给slave使用的密码
master_port: mysql-master的端口
master_log_file: master库执行SHOW MASTER STATUS;File的值
master_log_pos: master库执行SHOW MASTER STATUS;Position的值
```
```text
# 查看slave连接master的状态
SHOW SLAVE STATUS\G

当下面的Slave_IO_Running、Slave_SQL_Running为Yes时,证明slave可以同步master的数据了
若Slave_IO_Running为 No | Connecting,查看下面的 Last_IO_ERROR错误日志

*************************** 1. row ***************************
               Slave_IO_State: Waiting for master to send event
                  Master_Host: 192.168.33.10
                  Master_User: sync
                  Master_Port: 3306
                Connect_Retry: 30
              Master_Log_File: binlog.000003
          Read_Master_Log_Pos: 767
               Relay_Log_File: relay-bin.000002
                Relay_Log_Pos: 317
        Relay_Master_Log_File: binlog.000003
             Slave_IO_Running: Yes
            Slave_SQL_Running: Yes
              Replicate_Do_DB: 
          Replicate_Ignore_DB: 
           Replicate_Do_Table: 
       Replicate_Ignore_Table: 
      Replicate_Wild_Do_Table: 
  Replicate_Wild_Ignore_Table: 
                   Last_Errno: 0
                   Last_Error: 
                 Skip_Counter: 0
          Exec_Master_Log_Pos: 767
              Relay_Log_Space: 518
              Until_Condition: None
               Until_Log_File: 
                Until_Log_Pos: 0
           Master_SSL_Allowed: No
           Master_SSL_CA_File: 
           Master_SSL_CA_Path: 
              Master_SSL_Cert: 
            Master_SSL_Cipher: 
               Master_SSL_Key: 
        Seconds_Behind_Master: 0
Master_SSL_Verify_Server_Cert: No
                Last_IO_Errno: 0
                Last_IO_Error: 
               Last_SQL_Errno: 0
               Last_SQL_Error: 
  Replicate_Ignore_Server_Ids: 
             Master_Server_Id: 100
                  Master_UUID: adb44363-fc84-11ee-9daa-0242ac140228
             Master_Info_File: /var/lib/mysql/master.info
                    SQL_Delay: 0
          SQL_Remaining_Delay: NULL
      Slave_SQL_Running_State: Slave has read all relay log; waiting for more updates
           Master_Retry_Count: 86400
                  Master_Bind: 
      Last_IO_Error_Timestamp: 
     Last_SQL_Error_Timestamp: 
               Master_SSL_Crl: 
           Master_SSL_Crlpath: 
           Retrieved_Gtid_Set: 
            Executed_Gtid_Set: 
                Auto_Position: 0
         Replicate_Rewrite_DB: 
                 Channel_Name: 
           Master_TLS_Version: 
1 row in set (0.02 sec)
```

