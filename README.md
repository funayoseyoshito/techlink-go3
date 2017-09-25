# Go言語初心者向けハンズオン #3

## decker-compose

* 起動
  ```
  docker-compose up -d
  ```

* 起動中のコンテナに入る
  ```
  # app container
  docker-compose exec app bash
  
  # db container
  docker-compose exec db bash
  ```

* 停止
  ```
  docker-compose stop
  ```

* 停止(データも消える)
  ```
  docker-compose down
  ```
