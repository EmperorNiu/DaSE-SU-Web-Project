## Blog

### 看博客

- 从数据库中获取所有的blog信息, 得到blogs

  - ```json
    {
        "bolg":[
        {
            id:str,
            title:string,
            author:str,
            date:datatime,
            stars:int,
            views:int,
            favorites:int,
            recommends:str//(list[str]),
            content:text
            
        },
            
        {
        
    }
        ]
    }
    ```

- 遍历blogs, 循环渲染, 从每个blog中拿到title, author, date, stars, views, favorites, recommends , 作为卡片内容

  - 阅读按钮, 跳转到该blog内容界面(#TODO, 发送到blog页面渲染)
  - 稍后再看按钮, 记录一条数据, 存到数据库中, 传输值为blog_id

### 写博客

- 提交按钮
- 富文本编辑

### 稍后再看

同看博客部分

### 我的博客

同看博客部分, 区别在于获取到的blogs只包括自己的