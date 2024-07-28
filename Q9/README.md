# 練習問題9 成果物のジョブ間での受け渡し

`.github/workflows/release.yml` を編集して以下を全て満たすように作り変えてください。

- テスト/ビルドとリリースのジョブを分割する
    - ジョブ1
        - sample-appのテストとビルドを実行
        - ビルドされたバイナリファイル `sample-app-binary` をアップロード
    - ジョブ2
        - ジョブ1でビルドされた `sample-app-binary` をダウンロード
        - ビルドされたバイナリファイル `sample-app-binary` をリリースする

[回答例](./release.yml)

## 参考

- [ワークフロー内のジョブ間でのデータの受け渡し](https://docs.github.com/ja/enterprise-cloud@latest/actions/using-workflows/storing-workflow-data-as-artifacts#passing-data-between-jobs-in-a-workflow)
