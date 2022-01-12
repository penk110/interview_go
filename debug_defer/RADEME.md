### 问题

```
1.
func increaseA() int {
    var i int
    defer func() {
        i++
    }()
    return i
}

2.
func increaseB() (r int) {
    defer func() {
        r++
    }()
    return r
}
```

在不运行代码的情况下，能够全回答正确的还是少数（包括我自己）。有的同学可能知道答案，但不明白其中的原因。所以，有必要总结下，免得以后再掉坑里。试着先给出你的答案，再往下读！

### 什么是 defer

defer 是 Go 语言提供的一种用于注册延迟调用的机制，每一次 defer 都会把函数压入栈中，当前函数返回前再把延迟函数取出并执行。

> defer 语句并不会马上执行，而是会进入一个栈，函数 return 前，会按先进后出（FILO）的顺序执行。也就是说最先被定义的 defer 语句最后执行。**先进后出的原因是后面定义的函数可能会依赖前面的资源，自然要先执行；否则，如果前面先执行，那后面函数的依赖就没有了**。

### 采坑点

使用 defer 最容易采坑的地方是和**带命名返回参数的函数**一起使用时。

> defer 语句定义时，对外部变量的引用是有两种方式的，分别是作为**函数参数**和作为**闭包引用**。作为函数参数，则在 defer 定义时就把值传递给 defer，并被缓存起来；作为闭包引用的话，则会在 defer 函数真正调用时根据整个上下文确定当前的值。

避免掉坑的关键是要理解这条语句：

```
return xxx
```

这条语句并不是一个原子指令，经过编译之后，变成了三条指令：

```
1. 返回值 = xxx
2. 调用 defer 函数
3. 空的 return
```

1,3 步才是 return 语句真正的命令，**第 2 步是 defer 定义的语句，这里就有可能会操作返回值**。

我们一起来做几个题目：

```
1.
func f1() (r int) {
    defer func() {
        r++
    }()
    return 0
}

2.
func f2() (r int) {
    t := 5
    defer func() {
        t = t + 5
    }()
    return t
}

3.
func f3() (r int) {
    defer func(r int) {
        r = r + 5
    }(r)
    return 1
}
```

我们试着用上面的拆解方式，给出自己的答案，再往下看！
第一题拆解过程：

```
func f1() (r int) {

    // 1.赋值
    r = 0

    // 2.闭包引用，返回值被修改
    defer func() {
        r++
    }()

    // 3.空的 return
    return
}
```

**defer 是闭包引用，返回值被修改，所以 f() 返回 1。**

第二题拆解过程：

```
func f2() (r int) {
    t := 5
    // 1.赋值
    r = t

    // 2.闭包引用，但是没有修改返回值 r
    defer func() {
        t = t + 5
    }()

    // 3.空的 return
    return
}
```

**第二步没涉及返回值 r 的操作，所以返回 5。**

第三题拆解过程：

```
func f3() (r int) {

    // 1.赋值
    r = 1

    // 2.r 作为函数参数，不会修改要返回的那个 r 值
    defer func(r int) {
        r = r + 5
    }(r)

    // 3.空的 return
    return
}
```

**第二步，r 是作为函数参数使用，是一份复制，defer 语句里面的 r 和 外面的 r 其实是两个变量，里面变量的改变不会影响外层变量 r，所以不是返回 6 ，而是返回 1。**

文章写到这里，相信你可以很快解出文章开头的第二个例程，没错，第二个例程的 increaseB() 函数返回 1，分解如下：

```
func increaseB() (r int) {

    // 1.赋值
    r = 0

    // 2.闭包引用，r++
    defer func() {
        r++
    }()

    // 3.空 return
    return 
}
```

那第一个例程怎么解呢？我把代码再贴一遍：

```
func increaseA() int {
    var i int
    defer func() {
        i++
    }()
    return i
}
```

大家可能注意到，函数 increaseA() 是匿名返回值，返回局部变量，同时 defer 函数也会操作这个局部变量。对于匿名返回值来说，可以假定有一个变量存储返回值，比如假定返回值变量为 anony，上面的返回语句可以拆分成以下过程：

```
annoy = i
i++
return
```

由于 i 是整型，会将值拷贝给 anony，所以 defer 语句中修改 i 值，对函数返回值不造成影响，所以返回 0 。

### 后记

这篇文章算是一份学习总结，梳理下相关知识。看完文章之后，相信你一定掌握了这些细节，那就来试试今天的面试题吧，也是跟 defer 相关的。

最后，特别感谢 @饶大(公号：码农桃花源) ，相关阅读第一篇文章就是他的，非常棒，墙裂推荐！

相关阅读：
1.Golang之轻松化解defer的温柔陷阱https://segmentfault.com/a/1190000018169295#articleHeader4
2.Go defer实现原理剖析 https://my.oschina.net/renhc/blog/2870345
3.golang之defer简介 https://leokongwq.github.io/2016/10/15/golang-defer.html