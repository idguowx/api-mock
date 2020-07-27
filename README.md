# api-mock

**api mock tools**

http 接口模拟工具，提供以下功能：

1. 利用http接口设置，修改，查询请求与回复的映射
2. 自动序列化请求与回复的映射
3. 根据请求属性返回对应的回复



**编译与运行：**

1. **git** clone git@github.com:idguowx/api-mock.git

2. cd api-mock

3. go build 

4. cp -R api-mock data  /apps/api-mock

5. cd /apps/api-mock

6. ./api-mock -p 8099

   

### **概念定义**

**请求回复映射**

请求回复映射定义了请求与回复之间的关系。其属性包括以下：

testAppName：app名称，可以认为是命名空间。标识这个映射属于此命名空间。

caseName：用例名称，一个用例名称在app内唯一标识映射。

requestPattern：请求模式，定义了一个系列集合。只要符合此请求，都属于这个集合。

responseStr：回复字符串，当接收到一个属于requestPattern集合的请求，将会返回responseStr所设置的字符		串。



**requestPattern的定义**

格式：

```json
{
	"Method": {
		"Value": "Post",
		"Ignore": 2
	},
	"UriReg": {
		"Value": "/order/updateOrders",
		"Ignore": 2
	},
	"ContentTypeReg": {
		"Value": "^application/x-www-form-urlencoded$",
		"Ignore": 2
	},
	"BodyReg": {
		"Value": "^1=1$",
		"Ignore": 2
	}
}
```

Method：请求方法名称，如POST，GET

UriReg：请求Uri的正则描述，如UriReg设置为 "^/order/[a-z|A-Z]+$"，然后发起一个请求uri为 /order/find

​	则可以通过UriReg的检查。

ContentTypeReg：

BodyReg：请求体的正则描述。类似UriReg。

Ignore：是否忽略此字段，1是，2否。如果忽略就代表请求匹配的过程不会检查这个字段。



### **API**

1. **保存请求与回复映射**

   POST /matching/save

   参数：

   | 参数名         | 类型        | 备注         |
   | -------------- | ----------- | ------------ |
   | testAppName    | string      | app名称      |
   | caseName       | string      | 用例名称     |
   | requestPattern | json string | 请求描述模式 |
   | responseStr    | string      | 返回内容     |
   |                |             |              |

   返回：

   ```json
   {"Success":true,"ErrorCode":0,"ErrorMessage":"","Payload":null}
   ```

   

   

2. **删除请求与回复映射**

   POST /matching/del

   参数：

   | 参数名      | 类型   | 备注     |
   | ----------- | ------ | -------- |
   | testAppName | string | app名称  |
   | caseName    | string | 用例名称 |
   |             |        |          |

   返回：

   ```json
   {"Success":true,"ErrorCode":0,"ErrorMessage":"","Payload":null}
   ```

   

   

   ------

   

3.**查询请求与回复映射**

POST /matching/list

参数：

| 参数名      | 类型   | 备注     |
| ----------- | ------ | -------- |
| testAppName | string | 应用名称 |
|             |        |          |

返回：



```json
{
	"Success": true,
	"ErrorCode": 0,
	"ErrorMessage": "",
	"Payload": {
		"app1": {
			"case1": {
				"CaseName": "case1",
				"RequestPattern": {
					"Method": {
						"Value": "Post",
						"Ignore": 2
					},
					"UriReg": {
						"Value": "/order/updateOrders",
						"Ignore": 2
					},
					"ContentTypeReg": {
						"Value": "^application/x-www-form-urlencoded$",
						"Ignore": 2
					},
					"BodyReg": {
						"Value": "^1=1$",
						"Ignore": 2
					}
				},
				"ResponseStr": "22222"
			}
		},
		"app2": {
			"case1": {
				"CaseName": "case1",
				"RequestPattern": {
					"Method": {
						"Value": "Post",
						"Ignore": 2
					},
					"UriReg": {
						"Value": "/order/updateOrders",
						"Ignore": 2
					},
					"ContentTypeReg": {
						"Value": "^application/x-www-form-urlencoded$",
						"Ignore": 2
					},
					"BodyReg": {
						"Value": "^1=1$",
						"Ignore": 2
					}
				},
				"ResponseStr": "55555"
			}
		}
	}
}
```








