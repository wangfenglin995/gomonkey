# gomonkey
gomonkey是golang一个测试打桩框架，能够低成本完成打桩功能

[官网地址] https://github.com/agiledragon/gomonkey

[官方用例] https://github.com/fishingfly/gomonkey_examples

常用功能
1. 给函数打一个桩
2. 给函数打一个桩序列
3. 给成员函数打一个桩
4. 给成员函数打一个桩序列
5. 给函数变量打一个桩
6. 给全局变量打一个桩

其中用的比较多的是1和3两个场景

当我们要测试一个函数，但是函数会调用DAO获取数据库中、或者缓存里的数据， 
那么就可以把函数里调用DAO或者缓存数据的函数给mock掉，从而专注于我们业务逻辑的测试
