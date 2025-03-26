
# 游戏框架

### 代码框架由两个部分组成

#### mini_game
- 包括底层包 om_b、om_g、c_engine、om_struct 等逻辑代码;
- 并且提供模块包；

#### mini_admin
- 包括后台部分；
- 包括游戏业务部分, 主要是从 mini_game 包中调用;


### 部署框架由三个部分组成

#### mini_api
- 第三方接入
- 游戏业务处理
- 游戏逻辑处理

#### mini_cms
- 游戏进程处理
- 游戏逻辑验证

#### mini_admin
- 后台管理
- 游戏管理
