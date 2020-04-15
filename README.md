# Cat_wifi

现在大多数手机都有 `wifi` 分享功能, 但是返回的都是一个二维码, 所以我写一个一个解析器, 用来自动获取 `wifi` 的信息

## Usage

```bash

☁  cat_wifi [master] ⚡ go get -u github.com/d1y/cat_wifi

#-----

☁  cat_wifi [master] ⚡ go run .
2020/04/09 13:09:17 正在查找手机
2020/04/09 13:09:17 目前找到 1 台设备
2020/04/09 13:09:17 ===================
2020/04/09 13:09:17 该设备型号:  LLD_AL20
2020/04/09 13:09:17 自动连接该设备...
2020/04/09 13:09:17 开启屏幕中...
2020/04/09 13:09:19 手机已解锁
2020/04/09 13:09:19 自动打开Wifi设置
2020/04/09 13:09:20 将需要解密的wifi打开分享弹窗
我已打开分享弹窗[Y/N]: y
2020/04/09 13:10:05 输入正确
2020/04/09 13:10:06 正在截图, 临时文件:  /var/folders/fx/t5bzx1kx5kb4t44_d9hcsmy40000gn/T/screen.png
2020/04/09 13:10:06 Wifi名称:  陈哥哥的iPhone
2020/04/09 13:10:06 Wifi密码:  1008611
2020/04/09 13:10:06 ===================

```


## ADB

参考文章:
- [使用adb命令操控Android手机](https://www.jianshu.com/p/65e80c60f656)
- [awesome-adb](https://mazhuang.org/awesome-adb)

```
# 判断屏幕是否点亮
adb shell dumpsys display | grep "mScreenState"

# 判断屏幕是否锁定(密码锁)
adb shell dumpsys window | grep "mDreamingLockscreen"

# 点亮屏幕
adb shell input keyevent 26

# 启动 `wifi` 设置界面
adb shell am start -a android.intent.action.MAIN -n com.android.settings/.wifi.WifiSettings

# 截图
adb shell /system/bin/screencap -p /data/local/tmp/tmp.png

# 将文件拉取到本地
adb pull /data/local/tmp/tmp.png ~/Desktop/test.png

# 查看 `wifi` 密码(配置文件, 需要`root`)
cat /data/misc/wifi/*.conf

```

## TODO

- [ ] 考虑多台设备连接下, 设备选择
- [x] 设备已 `root` 下, 直接获取所有 `wifi` 密码