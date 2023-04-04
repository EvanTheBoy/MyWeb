### 结构

在项目伊始，我们需要弄清楚它的整体设计思路，这样才能在后续的开发过程中做到游刃有余，不至于乱了手脚。

既然要从零开始编写一个Web框架，那就必须要搞清楚Web server究竟是如何工作的，经过调查研究，我总结出了以下几点：

1. 通过创建一个Server数据结构来创建一个http服务；

2. Server通过for循环不断地监听来自客户端的请求；
3. 当监听到有一个请求的时候，就创建一个连接结构，并默认开启一个协程，即goroutine来为其服务；
4. 进行具体的业务逻辑处理；

整体的思路还是比较清晰的。我们可以先看看官方的Server数据结构长什么样：

```go
type Server struct {
    Addr string // 请求监听地址
    Handler Handler // 请求处理函数
}
```

当然，实际的Server结构远不止上面列出的那几个字段，这里只是为了方便目前的理解而省去了其他字段。

所以说，比较重要的是实现这个Handler处理函数。

### 实现

创建一个framework包，在这个包下建立一个core.go，在这个go文件中定义一个NewCore()方法，然后在main函数中直接调用这个方法，从而开始使用自己编写的Web框架，去进行一些常规的处理。

现在，思考一个问题，这个NewCore()函数到底是个什么东西？它的返回值又是什么呢？或者，这个方法需不需要返回值呢？其实NewCore()方法就是创建一个Core实例，然后把它返回出去，这样，外界拿到了这个core，就可以拿来做Web相关的一些操作了。不过目前因为Core里面还什么都没有，因此我们直接返回一个&Core{}。

然后，为了实现对外提供的ListenAndServer()方法，我们要在自己的Web框架中编写一个方法，来实现这个逻辑，这个方法还需要传入一些参数，让它可以处理写请求和读请求。

最后，代码写出来就是这样：

```go
type Core struct {
    
}

func NewCore() *Core {
    return &Core{}
}

func (c *Core) serveHTTP(request *http.Request, response http.ResponseWriter) {
    // To do
}
```

在main.go中：

```go
func main() {
    server := http.Server{
        Addr: ":8080",
        Handler: framework.NewCore(),
    }
    server.ListenAndServe()
}
```

那么在main函数中，我们自己创建了一个Server数据结构，自定义了监听地址，和Handler处理函数，从而实现了一个HTTP服务。
