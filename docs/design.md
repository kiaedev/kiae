ZOS, 为应用提供基础能力而生。一个应用运行平台。 

## 使用场景
1. 基于k3s构建边缘运行的应用平台
2. 基于k3s构建类似群辉DSM的应用平台
3. 基于k8s构建企业内部应用平台

参考群辉DSM的方式，平台的应用可以使用系统提供的账号权限等基础能力。可以直接在浏览器访问应用的地址，也可以以内置应用的方式访问。

## 系统架构

kui：kos的前端代码
kos：openkos的apiserver，提供api接口；实现内置应用及系统配置；提供基础能力API(通知、账户/组、Hook等)，管理自定义资源
runtime：一个基于kubebuilder构建的operator，实现Application的逻辑

kos-app-devops：一个基于kos的devops系统
- child-system1
- child-system2
- child-system3

## 核心：

安装逻辑：创建capp（由apiserver提供接口实现）：
1. 将demo.app.json转换成CloudAPP资源(如果依赖中间件则添加响应的traits)
2. apply CloudAPP资源，CloudAPP基于oam封装
3. capp有状态，刚安装完为NotReady
4. capp会自动处理依赖，比如自动申请数据库，kafka等操作
5. 依赖申请完毕后悔变为Ready状态，这时应用可以启动。

更新image版本可实现：发布/升级逻辑：

## 应用

### 系统应用
- Settings
- App Store
- ingress-controller
- cert-manager
- 集群管理：管控其他集群, infra一键安装等能力

### 普通应用
- Devops：提供多集群的构建、部署能力，底层仍然依赖kos-controller
- zpan
- rslocal
- autossl
- mysql
- mongodb
- for any self-hosted-app
- 

## 设置

### 系统设置
- Volumes
- SSL证书管理
- 账户管理

### 应用设置
- ingress
- config
- env
- secret
- 权限控制


