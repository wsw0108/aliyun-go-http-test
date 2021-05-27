# aliyun-go-http-test

## local
`http :9000/v1/headers x-fc-qualifier:v1_2_3`
`http :9000/headers x-fc-qualifier:v1_2_3`

## serverless

`serverless/aliyun`只支持创建事件函数（没法创建http触发器）。

## funcraft

`http http://custom-domain/headers`
`http http://custom-domain/v1/headers`
