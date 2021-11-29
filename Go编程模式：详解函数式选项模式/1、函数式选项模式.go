package main

import (
	"log"
)

//一般通过定义New函数来充当构造函数；对于结构有较多字段的，最好的方式是：函数式选项模式

func main()  {
	svr := New("localhost", 1234)
	if err := svr.Start(); err != nil {
		log.Fatal(err)
	}
}

//========结构示例=======

type Server struct {
	host string
	port int
}

func New(host string, port int) *Server  {
	return &Server{host, port}
}

func (s *Server)Start() error  {
	return nil
}



//但如果要扩展 Server 的配置选项，如何做？通常有三种做法：

//1、为每个不同的配置选项声明一个新的构造函数  Server 增加了两个字段；这种方式配置较少且不太会变化的情况
	//type Server {
	//	host string
	//	port int
	//	timeout time.Duration
	//	maxConn int
	//}
	//func New(host string, port int) *Server {
	//	return &Server{host, port, time.Minute, 100}
	//}
	//func NewWithTimeout(host string, port int, timeout time.Duration) *Server {
	//	return &Server{host, port, timeout}
	//}
	//
	//func NewWithTimeoutAndMaxConn(host string, port int, timeout time.Duration, maxConn int) *Server {
	//	return &Server{host, port, timeout, maxConn}
	//}

//2、定义一个新的 Config 结构体来保存配置信息；这种做法，即使将来增加更多配置选项，也可以轻松的完成扩展，不会破坏 Server 的 API
	//type Server {
	//	cfg Config
	//}
	//
	//type Config struct {
	//	Host string
	//	Port int
	//	Timeout time.Duration
	//	MaxConn int
	//}
	//
	//func New(cfg Config) *Server {
	//	return &Server{cfg}
	//}

//3、使用 Functional Option Pattern【函数式选项模式】；将来增加选项，只需要增加对应的 WithXXX 函数即可。
	//type Option func(*Server)
	//
	//func New(options ...Option) *Server {
	//	svr := &Server{}
	//	for _, f := range options {
	//		f(svr)
	//	}
	//	return svr
	//}
	//
	//func WithHost(host string) Option {
	//	return func(s *Server) {
	//		s.host = host
	//	}
	//}
	//
	//func WithPort(port int) Option {
	//	return func(s *Server) {
	//		s.port = port
	//	}
	//}
	//
	//func WithTimeout(timeout time.Duration) Option {
	//	return func(s *Server) {
	//		s.timeout = timeout
	//	}
	//}
	//
	//func WithMaxConn(maxConn int) Option {
	//	return func(s *Server) {
	//		s.maxConn = maxConn
	//	}
	//}

	//func main() {
	//	svr := New(
	//		WithHost("localhost"),
	//		WithPort(8080),
	//		WithTimeout(time.Minute),
	//		WithMaxConn(120),
	//	)
	//	if err := svr.Start(); err != nil {
	//		log.Fatal(err)
	//	}
	//}

	//Uber 的 Go 语言编程规范中提到该模式时，建议定义一个 Option 接口
	
	//type options struct {
	//	cache  bool
	//	logger *zap.Logger
	//}
	//
	//type Option interface {
	//	apply(*options)
	//}
	//
	//type cacheOption bool
	//
	//func (c cacheOption) apply(opts *options) {
	//	opts.cache = bool(c)
	//}
	//
	//func WithCache(c bool) Option {
	//	return cacheOption(c)
	//}
	//
	//type loggerOption struct {
	//	Log *zap.Logger
	//}
	//
	//func (l loggerOption) apply(opts *options) {
	//	opts.logger = l.Log
	//}
	//
	//func WithLogger(log *zap.Logger) Option {
	//	return loggerOption{Log: log}
	//}
	//
	//// Open creates a connection.
	//func Open(
	//	addr string,
	//	opts ...Option,
	//) (*Connection, error) {
	//	options := options{
	//		cache:  defaultCache,
	//		logger: zap.NewNop(),
	//	}
	//
	//	for _, o := range opts {
	//		o.apply(&options)
	//	}
	//
	//	// ...
	//}



