# 練習問題5

`.github/workflows/run-on-push.yml` を編集して以下を全て満たすように作り変えてください。

- ワークフローを次の2つのジョブに分割する
    - ジョブ1
        - 次の流れの処理を行う
            - ソースコードをチェクアウト
            - Goのバージョン1.22をインストール
            - sample-appのテストを実行
                - sample-appディレクトリ配下で `make test` を実行すればOK
    - ジョブ2
        - ジョブ1が成功したときだけ実行
        - 次の流れの処理を行う
            - ソースコードをチェクアウト
            - Goのバージョン1.22をインストール
            - sample-appのビルドを実行
                - sample-appディレクトリ配下で `make build` を実行すればOK

[回答例](./run-on-push.yml)

## 参考

- [ワークフローでジョブを使用する](https://docs.github.com/ja/actions/using-jobs/using-jobs-in-a-workflow)
- sample-app の `make test` を失敗させるには sample-app の main_test.go を以下のように変更すればOK 
    - `want := "run sample-app"` を任意の文字列に変更
        - 例) `want := "run sample-app-failure"`
