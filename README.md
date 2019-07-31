## 简介
企业微信消息推送

## 说明

### 动态添加job

postman

请求:post

url: http://localhost:8080/wx/robot/msg/add

参数格式：form-data

参数：

cron:0 0/30 7,8 2 8 ? 

key:xxxxx

msg:周三未激活的童鞋，下午13:00-15:30激活

status:1

msgType:markdown




### msgType
默认是：text并且支持@all
 
markdown支持颜色和样式但是不支持@all

### cron
quartz表达式

0 0/30 7,8 2 8 ?    8月2号7、8点每30分钟一次

0 25 09,20 ? * 1,2,3,4,5 每周一、二、三、四、五的9：25、20:25执行一次


### key
机器人的key

### msg
消息内容

### status
状态：0-无效，1-有效  默认0无效

### 执行逻辑
执行job，判断status==1执行，否则不执行


## 数据库

| 库名             | 库类型     | 说明         |
| ---------------- | ---------- | ------------ |
| xxx (schema:xxx) | PostgresSQL | 机器人消息数据库 |

> 表说明：

| 表名 |  说明  |
|---|---|
|WX_ROBOT_MSG   |   机器人消息表   |

> 字段说明：

| 字段名 | 类型  |  说明  |
|---|  ---|  ---|
|cron_tab   | varchar |   Cronb表达式  |
|msg   |text |   消息内容   |
|key   |varchar |   企业微信key   |
|status | integer   |   状态0-无效,1-有效   |
|msg_type | varchar   |   消息类型：text,markdown   |


> SQL

    create table wx_robot_msg
    (
     id serial not null
      constraint wx_robot_msg_pkey
       primary key,
     cron_tab varchar(100),
     msg text,
     key varchar(100),
     status integer default 0,
     msg_type varchar(20) default 'text'::character varying
    )
    ;
    
    create unique index wx_robot_msg_id_uindex
     on wx_robot_msg (id)
    ;
