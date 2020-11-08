## 博客首页

#### URL

GET /blog/first?page=pageNum

#### Request

##### Header:

| key   | 类型   | 描述               | 是否可为空 |
| ----- | ------ | ------------------ | ---------- |
| token | string | 登录产生的会话标识 | N          |

#### Response

##### Body:
```json
{
  "blogs": [
      {
          "id": 0,
          "title": "",
          "author": "",
          "create_time": "",
          "tags": [],
          "read_times": 0,
          "short_intro": ""  
      }
  ]
}
```

##### 属性说明：

| 变量名      | 类型   | 描述     | 是否可为空 |
| ----------- | ------ | -------- | ---------- |
| read_times  | int    | 阅读数   | N          |
| short_intro | string | 博文简介 | Y          |

## 博客正文

#### URL

GET /blog/article?id=blogId

#### Request

##### Header:

| key   | 类型   | 描述               | 是否可为空 |
| ----- | ------ | ------------------ | ---------- |
| token | string | 登录产生的会话标识 | N          |

#### Response

##### Body:
```json
{
    "id": 0,
    "title": "",
    "author": "",
    "create_time": "",
    "tags": [],
    "read_times": 0,
    "article_html": "",
    "avg_read_time": "",
    "anchors": [
    	{
            "id": 0,
            "position": ""
        }
    ]
}
```

##### 属性说明：

| 变量名        | 类型   | 描述                         | 是否可为空 |
| ------------- | ------ | ---------------------------- | ---------- |
| article_html  | str    | 博客正文                     | N          |
| avg_read_time | int    | 平均阅读时长(min)            | N          |
| anchors       | object | 具体属性待定，用于标注评论点 |            |

## 获取评论

#### URL

GET /blog/comment/read?id=blogId&anchor=anchorId

#### Request

##### Header:

| key   | 类型   | 描述               | 是否可为空 |
| ----- | ------ | ------------------ | ---------- |
| token | string | 登录产生的会话标识 | N          |

#### Response

##### Body:
```json
{
	"anchor_id": 0,
    "blod_id": 0,
    "position": "",
    "comments":[
        {
            "author": "",
    		"create_time": "",
            "content": ""
        }
    ]
}
```

##### 属性说明：
## 提交评论

#### URL

POST /blog/comment/write

#### Request

##### Header:

| key   | 类型   | 描述               | 是否可为空 |
| ----- | ------ | ------------------ | ---------- |
| token | string | 登录产生的会话标识 | N          |

##### Body:
```json
{
  "user_id": "buyer_id",
  "anchor": {
    "position": "",
    "comments":[
        {
            "content": ""
        }
    ]
}
}
```

#### Response

##### Body:
```json
{
	"anchor_id": 0
}
```

##### 属性说明：

