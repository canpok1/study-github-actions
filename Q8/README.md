# 練習問題8 タグpush時にGitHubへとリリースする

`.github/workflows/release.yml` を編集して以下を全て満たすように作り変えてください。

- 以下を行う
    - ソースコードをチェックアウト
    - Goのバージョン1.22をインストール
    - sample-appのテストを実行
        - sample-appディレクトリ配下で `make test` を実行すればOK
    - sample-appのビルドを実行
        - sample-appディレクトリ配下で `make build` を実行すればOK
    - ビルドされたバイナリファイル `sample-app-binary` をリリースする
        - 次のコマンドを実行すればOK
        - `gh release create {pushされたタグ名} sample-app/sample-app-binary --generate-notes --verify-tag`

[回答例](./run-on-push.yml)

## 参考

- [ワークフローで GitHub CLI を使用する](https://docs.github.com/ja/actions/using-workflows/using-github-cli-in-workflows)
- [gh release](https://cli.github.com/manual/gh_release)
- [既定の環境変数](https://docs.github.com/ja/actions/learn-github-actions/variables#default-environment-variables)
- [権限をジョブに割り当てる](https://docs.github.com/ja/enterprise-cloud@latest/actions/using-jobs/assigning-permissions-to-jobs)
