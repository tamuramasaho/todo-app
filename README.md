# todo-app
## 概要
タスクを管理するアプリケーションです。クリーンアーキテクチャを参考に作成しました。

## 使い方

$GOPATHは/Users/user/goのように設定します。
```
cd /Users/user/go/src/github.com/tamuramasaho
```

クローンします。
```
git clone git@github.com:tamuramasaho/tamuramasaho/todo-app.git
```
移動します。
```
cd todo-app
```

環境を作ります。
```
docker-compose build
docker-compose run --rm server sql-migrate up
docker-compose up
docker-compose down -vs
```

http://localhost:8000/todos

にアクセスします。

slackへの通知を使いたい場合はWEBHOOKURLを
slack APIからゲットして
/repository/todo.go
のconstを編集し、
main.goのcron関係のコードをアンコメントしてください。