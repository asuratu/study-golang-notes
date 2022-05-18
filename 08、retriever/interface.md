## 接口
```
type retriever interface {
    Get(string) string
}
func main() {
    var r  retriever
    r = getRetriever()
    bytes := r.Get("https://imooc.com")
}
```
## duck typing 鸭子类型
1. 像鸭子走路，像鸭子叫（长得像鸭子），那么就是鸭子
2. 描述事务的外部行为而非内部结构
3. 严格说go属于结构化类型系统，类似duck typing，没有动态绑定，而是编译时绑定

## 接口的定义
1. 接口由使用者定义
2. 接口的实现者只要实现接口里的方法
3. 接口的实现是隐式的

## 接口变量里有什么
1. 实现者的类型
2. 实现者的值
3. 接口变量自带指针
4. 接口变量同样采用值传递，几乎不需要使用接口的指针
5. 指针接收者只能以指针方式使用，值接收者两者都可以

## 查看接口变量
1. 表示任何和类型 interface{}
2. Type Assertion
3. Type Switch

