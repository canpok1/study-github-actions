# 練習問題10 ワークフローの再利用

`.github/workflows/run-on-push.yml` と `.github/workflows/release.yml` とで行っているテストとビルドのジョブを、再利用可能なワークフローとして `.github/workflows/test-and-build.yml` に作成してください。
その後、作成したワークフローを使うように `.github/workflows/run-on-push.yml` と `.github/workflows/release.yml` を修正してください。

## 回答例
- [run-on-push.yml](./run-on-push.yml)
- [release.yml](./release.yml)
- [test-and-build.yml](./test-and-build.yml)

## 参考

- [ワークフローの再利用](https://docs.github.com/ja/actions/using-workflows/reusing-workflows)
