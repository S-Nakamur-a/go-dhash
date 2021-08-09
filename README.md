# 概要

Goの練習のためdHashをGo実装してみた

## 動かし方

```sh
go run main.go <dHashを求めたい画像へのパス>
```

## 参考情報

dHashの正式実装を真面目に調べていないので以下の誤差がありえる

- Resize Algorithm: 本実装ではBilinearを採用
- Bit Assign: 本実装では8x8のバイナリに対し、left-topからみたビッグエンディアンで64bit整数に変更
