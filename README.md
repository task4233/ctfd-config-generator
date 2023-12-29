# ctfd-config-generator
[English Version is here.](./README-en.md)

ctfd-config-generatorはCTFdを用いたCTFの作問を支援するためのツールです。以下の特徴があります。

- [x] CTFdをCLIベースで管理する[ctfcli](https://github.com/CTFd/ctfcli)の設定ファイルの作成
- [x] 作問に利用するディレクトリやファイルの作成

## 新しい問題の作り方
以下のコマンドで環境を作ってください。実行には、[Go](https://go.dev/doc/install)とMakeが必要です。

```bash
make gen
```

実行が完了すると、以下のようなディレクトリが作成されます。

```bash
ジャンル名
└── 問題名
    ├── README.md       # 問題の概要を書いてください
    ├── build           # 問題サーバで実行されるファイルを配置してください(オプショナル)
    ├── challenge.yml   # CTFdの設定を書いてください
    ├── flag.txt        # Flagを書いてください
    ├── public          # 配布用ファイルを配置してください(オプショナル)
    ├── solver          # ソルバを配置してください(オプショナル)
    └── writeup
        └── README.md   # 作問者Writeupを配置してください
```

## コントリビュート
バグや要望などがあれば、Issueを作成するかPull Requestを作成してください。
