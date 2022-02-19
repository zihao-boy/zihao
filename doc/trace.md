## trace 协议

```bigquery
{
"traceId": "bd7a977555f6b982",
"name": "get-traces",
"id": "ebf33e1a81dc6f71",
"parentId": "bd7a977555f6b982",
"timestamp": 1458702548478000,
"duration": 354374,
"annotations": [
{
"endpoint": {
"serviceName": "zipkin-query",
"ipv4": "192.168.1.2",
"port": 9411
},
"timestamp": 1458702548786000,
"value": "cs"
}
],
"binaryAnnotations": [
{
"key": "lc",
"value": "JDBCSpanStore",
"endpoint": {
"serviceName": "zipkin-query",
"ipv4": "192.168.1.2",
"port": 9411
}
}
],
param:{
    reqHeader:'',
    resHeader:'',
    reqParam:'',
    resParam:''
}
}
```

```bigquery
traceId：标记一次请求的跟踪，相关的Spans都有相同的traceId；
id：span id；
name：span的名称，一般是接口方法的名称；
parentId：可选的id，当前Span的父Span id，通过parentId来保证Span之间的依赖关系，如果没有parentId，表示当前Span为根Span；
timestamp：Span创建时的时间戳，使用的单位是微秒（而不是毫秒），所有时间戳都有错误，包括主机之间的时钟偏差以及时间服务重新设置时钟的可能性，出于这个原因，Span应尽可能记录其duration；
duration：持续时间使用的单位是微秒（而不是毫秒）；
annotations：注释用于及时记录事件；有一组核心注释用于定义RPC请求的开始和结束；

cs:Client Send，客户端发起请求；
sr:Server Receive，服务器接受请求，开始处理；
ss:Server Send，服务器完成处理，给客户端应答；
cr:Client Receive，客户端接受应答从服务器；
binaryAnnotations：二进制注释，旨在提供有关RPC的额外信息
```
