## 模仿GitHub设计博客网站API

**沈方哲**  
**17343099**

****

### 简介

#### 1.API

API是 "Application Programming Interface"（应用程序接口）的缩写，用于描述一个类库的特征以及如何使用它。  
参考：[Github API v3](https://developer.github.com/v3/)

#### 2.REST

REST是 "Representational State Transfer"（表现层状态转移）的缩写，它是由罗伊·菲尔丁（Roy Fielding）在2000年提出的软件架构模式，用于描述创建HTTP API的标准方法。

REST是一种风格，而不是标准。因为既没有REST RFC，也没有 REST 协议规范或者类似的规定，它只是提供了一组设计原则和约束条件，主要用于客户端和服务器交互类的软件。Roy Fielding 把 REST 定义成一种架构风格，其目标是“使延迟和网络交互最小化，同时使组件实现的独立性和扩展性最大化”。

参考：[RESTful 基础](https://www.ics.uci.edu/~fielding/pubs/dissertation/rest_arch_style.htm)

### 设计

#### 1.REST API基本操作

* GET：用于获取资源，检索位于指定URL的资源的表示形式。  
* POST：用于在指定URL处新建资源。服务器为新资源分配URL，并将该URL返回给客户端。  
* PUT：用于更新资源。  
* DELETE：用于删除指定URL处的资源。  

#### 2.HTTP 状态码

* 200+：请求成功。  
* 300+：请求被重定向到另一个URL。  
* 400+：从客户端发起的错误已经发生。  
* 500+：从服务器发起的错误已经发生。  

#### 3.API 接口功能

本次作业的API设计只提供了博客网站最最最基本的功能：  
用户注册、登录、登出；创建、更新、删除博客；查询博客，查询某一博主的全部博客。

### 实现

见接口说明文档 api_design.yaml
