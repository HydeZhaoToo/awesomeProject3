

[TOC]

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	
	r.GET("/ping/:name", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": c.Param("name"),
		})
	})
	r.GET("/dsa",test)
	err := r.Run()
	if err != nil {
		return 
	}
}
func test(c *gin.Context)  {
	c.String(200,"ds'a'ds\n")
}
```

# 网有网规SAAS化初步摸底



1. 使用什么开发语言？是否使用框架 框架是什么？



2. 是否有第三方的集成  集成方式是什么 队列 http 其他?



3. 使用什么数据库 数据量多大 每天增长量多少?



4. 项目目前是否有鉴权认证系统?



5. 项目基础架构是什么 如何提供服务  B/S 还是C/S?



6. 项目部署在什么系统   arm还是X86? 系统版本是什么? 是否使用了容器化服务



7. 项目在SAAS后 计划部署在哪里  华为云  是否有出海计划?



8. 项目有几个环境 目前开发的流程是什么的?



9. 程序运行是否有状态  比如有状态文件生成 下次启动依赖?



10. 是否有文件存储功能 是否使用了对象存储？ 如 AWS S3 azure blob等 或者共享磁盘形式 如samba NFS 等。



11. 项目中是否有定时任务程序，或者一次性的任务程序 目前是如何实现的?



12. 项目的程序配置目前如何管理?



https://www.icloud.com.cn/keynote/0KebHlWpAu65JzGLP75hNh45g





