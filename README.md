# Simple-auth

シンプルな認証API

## 必要なもの

- Docker Compose

## インストール

1. `.env.sample`をコピーして`.env`を作成
2. `.env`の内容を修正
3. `docker-compose up -d`
4. DBに`app/docs/migration.sql`の内容を反映
5. `app/docs/insert_test_user.sql`を参考にユーザーを追加

## 使い方

1. `/v1/auth`でトークンを発行
2. `/v1/verify`でトークンの有効性をチェック