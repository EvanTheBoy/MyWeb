## 上下文超时控制

### Context

我们的框架需要一个上下文的超时控制逻辑。在Web server工作的主流程中我们知道，HTTP 服务会为每个请求创建一个 Goroutine 进行服务处理。但是这些服务处理，可能直接在本地服务，也有可能去下游服务获取数据。无论是在本地处理，还是去下游获取服务，都可能存在超时问题，进而导致HTTP服务不可用。

因此，我们需要设计这样一个功能，它具有超时控制能力，可以在树形逻辑链条中实现节点之间的**信息传递**和**共享**。为什么是“树形逻辑链条”呢？正如刚才所说，HTTP服务的处理可能直接在本地进行，也有可能要去下游取数据。这样就可以形成一个类似树形链的结构，故称为树形逻辑链条。

因此我们选择使用Context定时器为整个链条设置超时时间，时间一到，结束事件被触发，链条中正在处理的服务逻辑会监听到，从而结束整个逻辑链条，让后续操作不再进行。

虽然官方是有提供context标准库的，但我们要自己写一个!所以，我们先看看官方的context提供了些什么函数给我们，然后我们再用这些函数，实现自己的Context。

我们可以看到有下面三个比较常用的库函数：

```go
// 创建退出 Context
func WithCancel(parent Context) (ctx Context, cancel CancelFunc){}
// 创建有超时时间的 Context
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc){}
// 创建有截止时间的 Context
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc){}
```

WithCancel 直接创建可以操作退出的子节点，WithTimeout 为子节点设置了超时时间（还有多少时间结束），WithDeadline 为子节点设置了结束时间线（在什么时间结束）。看似功能不同，其实它们的核心内容是一样的。这是因为Context里面有两个核心的函数（已略去非相关内容）：

```go
type Context struct {
    Done() <-chan struct{}
}

type CancelFunc func()
```

这两个函数的作用是什么呢？在树形链上，一个节点同时充当着两个角色，一个是下游树的管理者，另一个就是上游树的被管理者。所以，一个节点就必须具备：

- 让整个下游树结束的能力，也就是CancelFunc
- 在上游树结束的时候被通知的能力，也就是Done方法。同时，因为通知是需要不断监听的，所以 Done() 方法需要通过 channel 作为返回值让使用方进行监听。

总之，CancelFunc是主动让下游结束，而Done是被上游通知结束。

[Go web 框架 gin 的使用 - 掘金 (juejin.cn)](https://juejin.cn/post/6957982755527344158)

关于为什么要使用IGroup这样一个接口，这是因为，如果哪一天，我们发现现有Group不满足我们新的需求，比如多了一些方法，或者现有方法需要改进，或者现有的某些方法不需要了，倘若没有这样一个接口，我们需要大改特改！既要改Group，也要修改使用了Group的地方。况且，如果我们决定抛弃使用现有Group，要编写一个新的Group来，那改动的地方会相当多，工作量也会陡增。因此我们使用接口，在core.Group里面，我们是这么实现的：

```go
func (c *Core) Group(prefix string) IGroup {
    return NewGroup(c, prefix)
}
```

倘若现有Group不要了，那我们直接修改这个Group函数的实现就好了，函数定义，返回值，参数都可以不用动，就用原来的就好；否则，改动的地方绝对会很多。