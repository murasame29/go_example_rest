# go_example_rest

| GET :/ping | pongを返す | o |
| --- | --- | --- |
| GET :/pong | pingを返す | o |
| GET :/users | 全てのユーザを返す | o |
| POST:/users  | ユーザを追加する | o |
| PUT : /users | 未実装 | o |
| DELETE:/users | 全てのユーザを削除する | o |
| GET: /users/:id | useridのユーザを取得 | x |
| POST:/users/:id | 未実装 | x |
| PUT:/users/:id | IDのユーザを編集 | x |
| DELETE/:users/:id | idのユーザを削除 | x |

細かい設定関連は省略しているため、各自実装は必須