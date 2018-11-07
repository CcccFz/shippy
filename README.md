consignment-service（货运服务）  
user-service（用户服务）   
vessel-service（货船服务）  
email-service（通知服务）  
shippy-ui（界面）  



go-micro自带API网关启动： `micro --registry=mdns api --handler=rpc --address=:8080 --namespace=shipy`   
如果使用默认的mdns作为服务发现，则运行每个服务和客户端启动时，需要加上`--registry=mdns`，如：`go run cli.go --registry=mdns`


作者： [Ewan Valentine](https://twitter.com/Ewan_Valentine)  
原文： [ewanvalentine.io](https://ewanvalentine.io/tag/go/)

相关翻译：[二向箔](https://blog.dingkewz.com/post/tech/go_ewan_microservices_in_golang_part_1/)、[wuYinIO](https://wuyin.io/2018/05/10/microservices-part-1-introduction-and-consignment-service/)
