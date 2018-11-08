简介：<br>
<br>
项目对iris框架进行重新分层与封装，将入口、配置、逻辑层、模型层、路由层区分开来；<br>
并引入了gorm，并做二次封装，让数据库分层和访问更加简单<br>
适用于使用go框架进行接口开发并希望简化目录结构的小伙伴，go如此简单~<br>
<br>
<br>
第一步：<br>
<br>
将GOPATH环境变量设置到当前目录<br>
<br>
项目目录<br>
<br>
src<br>
	——github.com		//库文件<br>
	——project1			//项目文件<br>
		——conf			//配置文件<br>
		——logic			//逻辑层<br>
		——modle			//模型层<br>
		——router		//路由层<br>
<br>
<br>
main.go		//项目入口文件<br>
<br>
<br>

go版本<br>
<br>
go version go1.11.1<br>
<br>
<br>


使用框架
iris:
https://github.com/kataras/iris

使用orm:
gorm
https://github.com/jinzhu/gorm

使用的mysql驱动：
go-sql-driver：
https://github.com/go-sql-driver/mysql

使用的redis客户端
https://github.com/go-redis/redis

使用的http请求库
https://github.com/imroc/req



