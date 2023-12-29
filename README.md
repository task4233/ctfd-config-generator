# ctfd-config-generator
[English Version is here.](./README-en.md)

ctfd-config-generatorはCTFdを用いたCTFの作問を支援するためのツールです。以下の特徴があります。

- [x] CTFdをCLIベースで管理する[ctfcli](https://github.com/CTFd/ctfcli)の設定ファイルの作成
- [x] 作問に利用するディレクトリやファイルの作成

## 新しい問題の作り方
以下の3ステップで行ってください。

### 1.mainブランチから新しいブランチを生やす
ブランチ名は `ジャンル名/問題名` にしてください。例えば、 `misc` ジャンルの `welcome` 問題を作る場合、 `misc/welcome` としてください。

```bash
git switch main
git pull origin main
git switch -c ${ジャンル名}/${問題名}
```

### 2.問題用の環境を作る
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

### 3.変更をcommitしてPushし、PRを作成する
まず、問題に必要なファイルをcommitしてpushしてください。  
その後、[PR](https://github.com/SecHack365-Fans/TsukuCTF2023/pulls) を作成してください。

- build ディレクトリには問題提供ファイルや docker 関連ファイルを置きます。特に Web 系の問題はここの中でプログラムを書きます。また、問題を作成するのに必要なファイルなどもここに置きます。
- public ディレクトリには**ユーザに渡すファイル**を置きます。OSINT や Forensics なら多くの場合，画像をここに置きます．このディレクトリとREADME.md の問題文・問題名のみがユーザに与えられます
- writeup ディレクトリには Writeup を置きます。フォルダ内に```README.md```というファイルを作成し、Markdown形式で Writeup を記載してください。画像等を埋め込むことも可能です。
- README.md ファイルには問題名、問題文を記述します(オプショナル)。

## コントリビュート
バグや要望などがあれば、Issueを作成するかPull Requestを作成してください。
