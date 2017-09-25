# Go言語初心者向けハンズオン #3

##  APIを動かす時



* create
  ```
  curl -F "name=test1" -F "email=test@test.com" localhost:4000/users
  curl -F "name=test2" -F "email=test@test.com" localhost:4000/users
  ```
* read
  ```
  curl localhost:4000/users/1
  ```

* update
  ```
  curl -F "name=test1change1" -X PUT -F "email=test@test.com" localhost:4000/users/1  
  ```

* delete
  ```
  curl -X DELETE localhost:4000/users/1
  ```

* index
  ```
  curl localhost:4000/users
  ```

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
