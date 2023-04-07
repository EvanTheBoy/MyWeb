[Go web 框架 gin 的使用 - 掘金 (juejin.cn)](https://juejin.cn/post/6957982755527344158)

关于为什么要使用IGroup这样一个接口，这是因为，如果哪一天，我们发现现有Group不满足我们新的需求，比如多了一些方法，或者现有方法需要改进，或者现有的某些方法不需要了，倘若没有这样一个接口，我们需要大改特改！既要改Group，也要修改使用了Group的地方。况且，如果我们决定抛弃使用现有Group，要编写一个新的Group来，那改动的地方会相当多，工作量也会陡增。因此我们使用接口，在core.Group里面，我们是这么实现的：

```go
func (c *Core) Group(prefix string) IGroup {
    return NewGroup(c, prefix)
}
```

倘若现有Group不要了，那我们直接修改这个Group函数的实现就好了，函数定义，返回值，参数都可以不用动，就用原来的就好；否则，改动的地方绝对会很多。