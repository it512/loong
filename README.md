# loong

``` logo
   __
  / /  _   _   _     _
 / /_,'o|,'o| / \/7,'o|
/___/|_,'|_,'/_n_/ |_,'
                   _//  (2024)
```
流程引擎

# 已经完成的功能  

## Activity  

1. service task
2. user task
3. start event
4. end event
5. link event (chrow and catch)
6. message intermediate throw event
7. exclusive gateway
8. sequence flow(defalut and condition)
9. boundary event(Error event only)

## 优先考虑内部ERP方面的需求
1. mongodb 移除，改用数据库
2. 内部代码优化，重点支持 exclusive gateway，user task，和 service task
3. 外部集成，和自动化部分也是重点，但会重新实现

## 距离v0.0.1 还需解决的问题

1. 并行网关还存在问题，目前暂不能正常工作(暂时搁置，优先考虑项目需求）
2. 包容网关没有实现
3. 常用的任务处理模式，例如转派等考虑实现


