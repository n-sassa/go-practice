# go-practice

~~go-swagger の練習~~  
↓  
現場の環境がある程度判明したので echo + sqlboiler + DDD もどきで再作成  
※DDD についてはキャッチアップする時間ないので雰囲気だけです。

# 動作確認

### 1.Docker の立ち上げ

```
docker-compose up -d --build
```

### 2.マイグレ

```
# docker内で実行
make migrate

or

# ローカルから実行
docker-compose exec app make migrate
```

### 3.サーバ接続

http://localhost:3000

# 主な技術

- GoLang
  - フレームワーク
    - echo
  - ORM
    - sqlboiler
  - その他
    - golang-migrate マイグレ
    - air ホットリロード
- その他
  - docker
  - make

# 注意点

windows の場合は docker が wsl を使用しないようにする。  
air のホットリロードが効かなくなるため。  
(windows home の場合は諦める)

docker の設定(右上の歯車アイコン)→General  
□ Use the WSL 2 based engine  
のチェックを外す
