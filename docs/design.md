ZOS, 为应用提供基础能力而生。一个应用运行平台。

参考群辉DSM的方式，平台的应用可以使用系统提供的账号权限等基础能力。可以直接在浏览器访问应用的地址，也可以以内置应用的方式访问。

## 核心：
启动app，拉起一个pod。

### CRD
- CloudAPP
cloudappfile：面向开发者的云应用描述文件，例：demo.caf

- apiserver: 提供基础能力API(通知、账户/组、Hook等)，管理自定义资源
- controller-manger: 提供自定义资源到k8s原生资源的转换能力

### 基础能力
- 账号，权限等等

### 可以想象的应用
- for any self-hosted-app
- autossl
