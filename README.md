<h1 align=center>Chat-Skn</h1>

## WebSocket

- 心跳机制

## 用户

## 服务(Guild)/频道

- 成员
- 邀请
- 权限

## 消息

- 在线发送
  - 私聊
    - 使用ws直接推送
  - 频道
    - 获取
- 离线发送
  - 缓存在Redis中
- 提取并归类
  - 媒体
  - 链接
  - 文件

```json
// unread:{user_id}

{
    "unread:{user_id}": {
        "count": 0,
        "private": {
            "{{receiver_id}}": {
                "count": 0,
                "last_msg_id": "{{id}}"
            }
        },
        "channel": {
            "{{channel_id}}": {
                "count": 0,
                "last_msg_id": "{{id}}"
            }
        }
    }
}
```
## 搜索(elasticsearch)

- 频道
- 用户
- 频道内的消息

## 插件

### 答题

- OCR
- 提取出问题、答案
- 搜索问题
- 模拟练习

### 问答
