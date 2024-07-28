# 練習問題4 ソースコードのテストとビルド

`.github/workflows/run-on-push.yml` を編集して以下を全て満たすように作り変えてください。

- ワークフローでは以下を実行する
    - ソースコードをチェクアウト
    - Goのバージョン1.22をインストール
    - sample-appのテストを実行
        - sample-appディレクトリ配下で `make test` を実行すればOK
    - sample-appのビルドを実行
        - sample-appディレクトリ配下で `make build` を実行すればOK

[回答例](./run-on-push.yml)

## 参考

- [Goでのビルドとテスト](https://docs.github.com/ja/actions/automating-builds-and-tests/building-and-testing-go)
- [setup-go](https://github.com/actions/setup-go/)
