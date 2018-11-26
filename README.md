雑にgo言語でtodoアプリ作りました
各種インストールやデータベースの設定等は適当に変えてリポジトリにあげてるので設定が必要(メモ)

# 環境構築
## Goのインストール
brew install go

## パスを通す
bashrcかどこかに追加
export GOPATH=$HOME/(作業ディレクトリ)
export PATH=$PATH:$GOPATH/bin

## 必要なライブラリ追加
go get golang.org/x/tools/cmd/goimports
go get github.com/gin-gonic/gin

# 各種コマンド等
## サーバ起動
go run main.go
