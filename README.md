
## 游戏框架

#### 代码框架由两个部分组成

- mini_game
    1.  包括底层包 om_b、om_g、c_engine、om_struct 等逻辑代码;
    2.  并且提供模块包；

- mini_admin
    1.  包括后台部分；
    2.  包括游戏业务部分, 主要是从 mini_game 包中调用;


#### 部署框架由三个部分组成

- mini_api
    1.  第三方接入
    2.  游戏业务处理
    3.  游戏逻辑处理

- mini_cms
    1.  游戏进程处理
    2.  游戏逻辑验证

- mini_admin
    1.  后台管理
    2.  游戏管理