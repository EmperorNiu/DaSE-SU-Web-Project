## 项目展示首页

#### URL

GET /api/getprojlist

#### Request

##### Header:

| key   | 类型   | 描述               | 是否可为空 |
| ----- | ------ | ------------------ | ---------- |
| token | string | 登录产生的会话标识 | N          |

#### Response

##### Body:

```json
{
    "status":0,
  "message": [
      {
          "id": 0,
          "title": "",
          "author": "",
          "create_time": "",
          "img_url":""
      }
  ]
}
```

## 项目详情

#### URL

GET /api/getprojlist/:projid

#### Request

##### Header:

| key   | 类型   | 描述               | 是否可为空 |
| ----- | ------ | ------------------ | ---------- |
| token | string | 登录产生的会话标识 | N          |

#### Response

##### Body:
```json
{
    "status":0,
    "message":[{
    "id": 0,
    "title": "",
    "author": "",
    "create_time": "",
    "article_html": ""
    }]
}
```

##### 属性说明：

| 变量名       | 类型 | 描述 | 是否可为空 |
| ------------ | ---- | ---- | ---------- |
| article_html | str  | 正文 | N          |
|              |      |      |            |
|              |      |      |            |

## 获取评论

#### URL

GET /api/getcomments/:artid?pageindex=1

#### Request

##### Header:

| key   | 类型   | 描述               | 是否可为空 |
| ----- | ------ | ------------------ | ---------- |
| token | string | 登录产生的会话标识 | N          |

#### Response

##### Body:
```json
{
    "status":0,
    "message":[
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

POST /api/postcomment/:artid

#### Request

##### Header:

| key   | 类型   | 描述               | 是否可为空 |
| ----- | ------ | ------------------ | ---------- |
| token | string | 登录产生的会话标识 | N          |

##### Body:
```json
{
    "artid":0,
    "comments":[
        {
            "content": ""
        }
  ]
}
```

#### Response

##### Body:

```json
{
    "status":0,
	"message": "评论提交成功"
}
```

##### 属性说明：

## 校友会展示首页

#### URL

GET /api/getmemlist

#### Request

##### Header:

| key   | 类型   | 描述               | 是否可为空 |
| ----- | ------ | ------------------ | ---------- |
| token | string | 登录产生的会话标识 | N          |

#### Response

##### Body:

```json
{
    "status":0,
    "message": [
      {
          "id": 0,
          "author": "",
          "img_url":""
      }
  ]
}
```

## 校友信息

#### URL

GET /api/getmemlist/:memid

#### Request

##### Header:

| key   | 类型   | 描述               | 是否可为空 |
| ----- | ------ | ------------------ | ---------- |
| token | string | 登录产生的会话标识 | N          |

#### Response

##### Body:

```json
{
    "status":0,
    "message":[{
    "id": 0,
    "author": "",
    "article_html": ""
    }]
}
```

##### 属性说明：

| 变量名       | 类型 | 描述 | 是否可为空 |
| ------------ | ---- | ---- | ---------- |
| article_html | str  | 正文 | N          |
|              |      |      |            |