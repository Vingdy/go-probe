### adn_tracking version check

go版本要求：`1.11` 及以上

钉钉机器人示例

###adn_tracking版本
####服务器地区：sg   总数量:21
>版本号:version|3.19.4-1b881a9a 数量：15 版本号:version|3.19.3-3093470f 数量：6

配置修改：  

dingding/dingding.go  
`access_token = <钉钉机器人token>`  

version_check.go  
`Port = <开启端口号>//默认为11111`  

watcher/watcher.go  
`//过滤设置  
    region_filter = []string{   
 		"vg",  
 		"sg",  
 		"fk",  
 	}  
 	ad_type_filter = []string{  
 		"adn",
 	}
 	server_filter = []string{
 		"tktracking",
 		"tktrackingStaticPostback",
 		"tktrackingForRta",
 		"tktrackingInstallService",
 	}`
  
quick start:  
  
安装go环境  

在上层目录make get安装依赖

make build

make run

等待数秒后搜索完集群后curl 127.0.0.1:11111  

其他办法：  
  
摸索中  