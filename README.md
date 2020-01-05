# gorm_dao_reuse
分层情况下，使service共用dao层代码，更换orm 只需替换实现即可

此文件为了解决， 分层情况下使多个service公用一个dao
例如： 添加用户时需要进行 insert(user), 添加其他时需要写很多 insert(other)
find , delete , update 同理，在service层调用这些公用的方法 不用重复去写dao逻辑， 减少重复代码
使用这些方法需要在 service层 来构建所需要的model结构， 来满足增删改查需要
当然如果准备直接在 service层写对数据库的增删改查代码，请忽略此文件
因为相对来讲orm封装本身就够简单，此文件解决的是（分层下）service去调用dao的时候，减少dao的代码量
