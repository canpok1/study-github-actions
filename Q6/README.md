# 練習問題6 タグpush時に動くワークフローの追加

`.github/workflows/release.yml` を追加して以下を全て満たすワークフローを作成してください。

- ワークフローの名前は `Release` に設定する
- ワークフローでは `echo "release"` を実行する
- リポジトリに以下の名前のタグをpushしたときに動くよう設定する
    - `vX.X.X` (Xは任意の数字)
        - 例) `v1.23.45`

[回答例](./release.yml)

## 参考

- [特定のタグのプッシュが発生した場合にのみワークフローを実行する](https://docs.github.com/ja/actions/using-workflows/events-that-trigger-workflows#running-your-workflow-only-when-a-push-of-specific-tags-occurs)
