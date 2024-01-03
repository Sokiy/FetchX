# FetchX
Fetch Zinc Data Script Write By Go


## 使用方法

从 [Release](https://github.com/Sokiy/FetchX/releases) 页面下载对应平台的可执行文件，目前仅仅支持 Windows。（Linux 可自行编译）  
然后执行可以查看对应的帮助信息。
```shell
gougou.exe -h
gougou.exe --help
```
执行两个命令 fetch 和 version

### gougou fetch
```shell
gougou fetch
```
则会按照 config.toml 的配置进行下载

config.toml 配置文件说明
```toml
# 配置文件参数说明
# url: 下载地址
download_url = "http://zinc15.docking.org/substances/download/"
# total: 下载总数
total = 500000
# per_page: 每页下载数量
per_page = 100
# start_page: 开始页数
start_page = 1
# download_dir_name: 下载目录名称
download_dir_name = "Download"
# download_sub_dir_name: 下载子目录名称
download_sub_dir_name = "Scaffolds"
# file_prefix: 文件前缀
file_prefix = "ZINC_SCAFFOLDS_"
# file_suffix: 文件后缀
file_suffix = "sdf"
```


### gougou version
```shell
gougou version
```
则会显示当前版本信息


## 版本介绍

```text
v0.0.1 基础下载功能完善
v0.0.2 增加支持定义文件后缀 file_suffix
```
