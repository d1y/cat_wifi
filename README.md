# Cat_wifi

现在大多数手机都有 `wifi` 分享功能, 但是返回的都是一个二维码, 所以我写一个一个解析器, 用来自动获取 `wifi` 的信息

## Usage

先编译出来你当前系统的一个版本

```bash
git clone https://github.com/d1y/cat_wifi catf
cd catf
go build .

#-----

catf "WIFI:T:WPA;S:FAST_TEST;P:6666;;"

wifi加密类型:  WPA
wifi名称:  FAST_TEST
wifi密码:  6666

```

## TODO

- [ ] 读取二维码图片自动识别