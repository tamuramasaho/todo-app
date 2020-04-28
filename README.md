# todo-app
## 概要
タスクを管理するアプリケーションです。クリーンアーキテクチャを参考に作成しました。

## 使い方

$GOPATHは/Users/user/goのように設定します。
```
cd /Users/user/go/src
```

クローンします。
```
git clone git@github.com:tamuramasaho/todo-app.git
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