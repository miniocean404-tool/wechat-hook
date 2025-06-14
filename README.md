# wxhelper 使用

1. 文章网址：https://www.uniqchat.com/58.html
2. wxhelper 破解 hook dll ：https://github.com/ttttupup/wxhelper
3. 注入工具：https://github.com/nefarius/Injector
4. 注入工具 2：https://github.com/winsiderss/systeminformer、https://processhacker.sourceforge.io/ 未使用过
5. 微信历史版本：https://github.com/tom-snow/wechat-windows-versions/releases

# 基本原理

启动指定版本 PC 微信以后，利用注入程序将 dll 文件注入到微信进程内，可以截获所有的新消息，传递给外部接口，并提供发送消息的端口供外部程序调用。

# 安装微信

请安装资源包中提供的指定版本微信，历史版本中下载 3.9.8.25

安装完成后，请图标上右键【以管理员身份运行】启动微信，正常扫码登录微信

# 注入微信进程

1. 以管理员身份运行微信

2. 点击电脑左下角启动图标，搜索【cmd】，点击【以管理员身份运行】终端

切换到程序包目录下，比如程序包目录在 D:\software

先切换到 D 盘，输入 D:

```
C:\Windows\System32>D:
```

再切换到 software ，输入 cd software

```
D:\>cd software
```

3. 粘贴并运行以下命令(自己的联想电脑是使用 64 位 Injector.exe)

```
Injector.exe -n WeChat.exe -i wxhelper3.9.8.25.dll
```

出现：successfuly injected module 表示成功

显示上面的提示文字，代表注入成功！否则，都是注入失败

以上操作完成后，微信进程会提供出收发消息接口

自行根据下面文档，开发监听程序收取消息，处理消息和发送消息

# 接口文档

https://github.com/ttttupup/wxhelper/blob/dev-3.9.8.25/doc/3.9.8.25.md
