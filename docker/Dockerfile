# 【参考】Go アプリケーションのビルド、テスト、AWS ECS へのデプロイを行う方法
# https://circleci.com/ja/blog/use-circleci-orbs-to-build-test-and-deploy-a-simple-go-application-to-aws-ecs/
# https://gist.github.com/daumie/43317089caa48938fd6dab4fa2bc0193

# ビルド用ステージ
FROM golang:1.21-alpine as builder

WORKDIR /build

COPY ./ /build/

RUN ls -altrh

RUN go build -o /build/main /build/cmd/main.go


# 実行用ステージ
FROM alpine:3.18.3

WORKDIR /app

# 【参考】Alpine Linux でタイムゾーンを変更する
# https://qiita.com/dtan4/items/8359e389b95cbc60952d
# 記事では削除も行っているが、行うとコンテナ実行時にエラーとなったため、行わないようにした
RUN apk --update add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime

# rootではない実行用ユーザーを作成
RUN adduser -S -D -H -h /app appuser
USER appuser

COPY --from=builder /build/main /app/
# .envファイルがないとエラーになるため、ファイルを作成。環境変数の値はECSのタスク定義にて設定されたものが使用される
COPY --from=builder /build/.env /app/.env
# メール本文が格納されているディレクトリをコピー
# COPY --from=builder /build/resources /app/resources

ENTRYPOINT ["./main"]