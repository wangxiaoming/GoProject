## Go Log包使用
log包定义了Logger类型，该类型提供了一些格式化输出的方法。
本包也提供了一个预定义的“标准”logger，可以通过调用函数Print系列(Print|Printf|Println）、Fatal系列（Fatal|Fatalf|Fatalln）、和Panic系列（Panic|Panicf|Panicln）来使用，比自行创建一个logger对象更容易使用。

```
package main

import (
	"log"
)

func main() {
	log.Println("这是一条很普通的日志。")
	v := "很普通的"
	log.Printf("这是一条%s日志。\n", v)
	log.Fatalln("这是一条会触发fatal的日志。")
	log.Panicln("这是一条会触发panic的日志。")
}
```
运行结果
```
2021/03/03 14:38:33 这是一条很普通的日志。
2021/03/03 14:38:33 这是一条很普通的日志。
2021/03/03 14:38:33 这是一条会触发fatal的日志。

Process finished with exit code 1
```
logger会打印每条日志信息的日期、时间，默认输出到系统的标准错误。Fatal系列函数会在写入日志信息后调用os.Exit(1)。Panic系列函数会在写入日志信息后panic。
### 配置logger

默认情况下的logger只会提供日志的时间信息，但是很多情况下我们希望得到更多信息，比如记录该日志的文件名和行号等。log标准库中为我们提供了定制这些设置的方法。

log标准库中的Flags函数会返回标准logger的输出配置，而SetFlags函数用来设置标准logger的输出配置。

```
func Flags() int
func SetFlags(flag int)
```

#### flag 选项
```
const (
    // 控制输出日志信息的细节，不能控制输出的顺序和格式。
    // 输出的日志在每一项后会有一个冒号分隔：例如2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
    Ldate         = 1 << iota     // 日期：2009/01/23
    Ltime                         // 时间：01:23:23
    Lmicroseconds                 // 微秒级别的时间：01:23:23.123123（用于增强Ltime位）
    Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
    Lshortfile                    // 文件名+行号：d.go:23（会覆盖掉Llongfile）
    LUTC                          // 使用UTC时间
    LstdFlags     = Ldate | Ltime // 标准logger的初始值
)
```

#### 测试代码
```
func main() {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
}
```
运行结果
```
2021/03/03 15:05:25.961952 D:/GoProject/src/main/gobase/log/logger1.go:7: 这是一条很普通的日志。
```

### 配置输出位置
```
func SetOutput(w io.Writer)
```
SetOutput函数用来设置标准logger的输出目的地，默认是标准错误输出。

例如，下面的代码会把日志输出到同目录下的xx.log文件中。

```
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	logFile, err := os.OpenFile("D:/logs/xxxx/Payment/test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
	log.SetPrefix("[小王子]")
	log.Println("这是一条很普通的日志。")
}
```

运行结果：D:/logs/xxxx/Payment/test.log
```
2021/03/03 15:16:19.885118 D:/GoProject/src/main/gobase/log/loogfile.go:17: 这是一条很普通的日志。
[小王子]2021/03/03 15:16:19.921021 D:/GoProject/src/main/gobase/log/loogfile.go:19: 这是一条很普通的日志。
```
### 创建 logger 
log标准库中还提供了一个创建新logger对象的构造函数–New，支持我们创建自己的logger示例。New函数的签名如下：

```
func New(out io.Writer, prefix string, flag int) *Logger
```
New创建一个Logger对象。其中，参数out设置日志信息写入的目的地。参数prefix会添加到生成的每一条日志前面。参数flag定义日志的属性（时间、文件等等）。

```
func main() {
	logger := log.New(os.Stdout, "<New>", log.Lshortfile|log.Ldate|log.Ltime)
	logger.Println("这是自定义的logger记录的日志。")
}
```

运行结果
```
<New>2017/06/19 14:06:51 main.go:34: 这是自定义的logger记录的日志。
```

### 配置日志前缀
log标准库中还提供了关于日志信息前缀的两个方法：

```
func Prefix() string
func SetPrefix(prefix string)
```
其中Prefix函数用来查看标准logger的输出前缀，SetPrefix函数用来设置输出前缀。

```
func main() {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
	log.SetPrefix("[小王子]")
	log.Println("这是一条很普通的日志。")
}
```
上面的代码输出如下：
```
[小王子]2017/06/19 14:05:57.940542 .../log_demo/main.go:13: 这是一条很普通的日志
```
这样我们就能够在代码中为我们的日志信息添加指定的前缀，方便之后对日志信息进行检索和处理

##### 欢迎关注公众号：程序员开发者社区
![在这里插入图片描述](https://img-blog.csdnimg.cn/20210303152343266.jpg)

##### 参考资料
- https://www.liwenzhou.com/posts/Go/go_log/
