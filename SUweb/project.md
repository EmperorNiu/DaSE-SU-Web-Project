## 项目展示首页

baseUrl: http://47.116.103.155:8001/api/project

#### URL

GET /getProjectList

#### Request

##### Header:

| key   | 类型   | 描述               | 是否可为空 |
| ----- | ------ | ------------------ | ---------- |
| token | string | 登录产生的会话标识 | N          |

#### Response

##### Body:

```json
{
    "message": "string",
    "projects": [
      {
          "project_id": 2,
          "project_name": "testProject",
          "project_leader": "testLeader",
          "ProjectDescription": "",
          "Members": null,
          "Tutor": {
              "TutorId": 0,
              "TutorName": ""
          },
          "TutorId": 0,
          "Comments": null
      }
  ]
}
```

## 项目详情

#### URL

GET /getprojlist?id=0

#### Request

##### Header:

| key   | 类型   | 描述               | 是否可为空 |
| ----- | ------ | ------------------ | ---------- |
| token | string | 登录产生的会话标识 | N          |

##### Params

| key   | 类型   | 描述               | 是否可为空 |
| ----- | ------ | ------------------ | ---------- |
| id | string | 项目id | Y 默认为0          |

#### Response

##### Body:
```json
{
    "message": "success",
    "project": {
        "project_id": 2,
        "project_name": "testProject",
        "project_leader": "testLeader",
        "ProjectDescription": "",
        "Members": null,
        "Tutor": {
            "TutorId": 0,
            "TutorName": ""
        },
        "TutorId": 0,
        "Comments": null
    }
}
```

##### 属性说明：

| 变量名       | 类型 | 描述 | 是否可为空 |
| ------------ | ---- | ---- | ---------- |
| Null | str  | 正文 | N          |
|              |      |      |            |
|              |      |      |            |

## 获取评论

#### URL

GET /getcomments?project_id=0

#### Request

##### Header:

| key   | 类型   | 描述               | 是否可为空 |
| ----- | ------ | ------------------ | ---------- |
| token | string | 登录产生的会话标识 | N          |

##### Params

| key   | 类型   | 描述               | 是否可为空 |
| ----- | ------ | ------------------ | ---------- |
| project_id | string | 项目id | N          |

#### Response

##### Body:
```json
{
    "message": "string",
    "comments":[
        {
            "comment_id": "",
    		"created_at": "time",
            "auth_id": "",
            "content": ""
        }
    ]
}
```

##### 属性说明：

## 提交评论

#### URL

POST /comment

#### Request

##### Header:

| key   | 类型   | 描述               | 是否可为空 |
| ----- | ------ | ------------------ | ---------- |
| token | string | 登录产生的会话标识 | N          |

##### Body:
```json
{
    "artid":0,
    "comment":{
          "auth_id": "",
          "content": ""
     }
}
```

#### Response

##### Body:

```json
{
	"message": "string"
}
```

##### 属性说明：

## 校友会展示首页

#### URL

GET /getMemList

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