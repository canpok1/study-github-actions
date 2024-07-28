# 練習問題11 コンテナでのジョブ実行

`.github/workflows/test-and-build.yml` を以下のように修正してください。

- `golang:1.22` のコンテナイメージをジョブで使うように設定
- Goのインストールを行うstepを削除

## 回答例
- [test-and-build.yml](./test-and-build.yml)

## 参考

- [コンテナ内でのジョブの実行](https://docs.github.com/ja/actions/using-jobs/running-jobs-in-a-container)
