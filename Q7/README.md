# 練習問題7 特定のブランチpush時に動くようにする

`.github/workflows/run-on-push.yml` を編集して以下を全て満たすように作り変えてください。

- リポジトリに次のいずれかのブランチがpushされた時だけ動くようにする
    - `feature/*` (`*`は任意の文字列)
    - `main`

[回答例](./run-on-push.yml)

## 参考

[特定のブランチへのプッシュが発生した場合にのみワークフローを実行する](https://docs.github.com/ja/actions/using-workflows/events-that-trigger-workflows#running-your-workflow-only-when-a-push-to-specific-branches-occurs)
