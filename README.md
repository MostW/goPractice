# goPractice

趁着假期学习go 的使用，打算从模仿别人的项目入手，因此参考了[**FIFA-World-Cup**](https://github.com/GopherCoder/FIFA-World-Cup)，用到的第三方库为gin和gorm，数据库使用mysql。目前实现的功能：

1.登录注册：在请求头加上Authorization，带上token的方式实现。

2.待续。

在练习的过程中解决的一些问题：

1.原项目的"golang.org/x/crypto/bcrypt"包无法引入，导致密码加密和生成token的问题

2.原项目中使用struct的标签导致无法生成表

3.原项目是直接postman请求接口，因此没有跨域问题，但是与前端结合起来的时候，在浏览器中有跨域问题。

4.gin中间件的使用中遇到的一些问题

5.前端部署。
