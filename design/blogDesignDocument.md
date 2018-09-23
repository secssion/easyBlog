# 我的博客系统

## 登录功能的设置

### 字段设置

| 字段 | 备注 |
| - | - |
| username | 用户账号 |
| password | 用户密码 |
| identifying | 邮箱验证码， 登录的时候，会收到一条邮箱验证码,(邮箱为注册的时候指定的账号密码) |
* 是否发送验证码，前端界面给一个图片验证，前端验证成功之后才可以，请求邮箱服务器给用户发送验证码
* 双重验证是为了保证用户数据的安全，前端验证是为了保证每次请求都是 ‘人’ 在操作， 后端验证是为了保证用户正确的拿到自己的数据

## 用户表单的设置

>* 这个设计用于注册账号密码（用户注册功能不对外开放）

### 字段设置
|字段名| 备注|
|-|-|
|user_id | 需要保证唯一性|
|username | 用户账号 |
|password | 用户密码 |
|email | 用户邮箱 |
| telephone | 用户手机|
| nickname | 用户别称|
| register_time | 用户注册时间，可以用于显示用户使用博客多久
| head_icon | 用户头像的icon |

> * 如果觉得有必要新增加字段，可以讨论下

## 密码的加密处理
* 所有设计密码相关的的 功能需求都应该对其进行加密处理
* 加密处理，前后端都采用 MD5的加密方式处理,或者选择RSA的非对称加密算法进行解析