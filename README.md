# study-go-mercaridoc
go学習のアウトプット

メルカリの[プログラミング言語Go完全入門](https://drive.google.com/file/d/1fLlg3Xw7CV680GQ65WkjxU5qX-PsApJg/view)をベースにした個人学習記録です.

コーディング規約の参考：[uver-go](https://github.com/uber-go/guide) (未読)

|章|お題|作成日|習得要素|備考
:--|:--|:--|:--|:--
|基本構文|奇数と偶数|2020/11/08|制御構文|
|基本構文|おみくじプログラムを作ろう|2020/11/09|fmt.Scan/rand.Seed|
|関数と型|スライスの利用|2020/11/09|スライス|
|関数と型|ユーザー定義型の使用|2020/11/10|type,map,struct
|関数と型|奇数偶数判定関数|2020/11/11|関数
|関数と型|複数戻り値の利用|2020/11/12|複数戻り値
|関数と型|ポインタ|2020/11/12|ポインタ
|関数と型|レシーバに変更を加える|2020/11/12|レシーバ
|パッケージ|パッケージを分けてみよう|2020/11/13|パッケージ
|パッケージ|ライブラリを取得してみよう|2020/11/16|go get
|パッケージ|modulesを使ってみよう|2020/11/16|go get
|コマンドラインツール|プログラム引数|2020/11/17|引数
|コマンドラインツール|flagパッケージ|2020/11/17|flag
|コマンドラインツール|入出力|2020/11/20|ファイル処理
|コマンドラインツール|catコマンドを作ろう|2020/11/22|scanner,byte.buf,Fprint
|コマンドラインツール|画像変換コマンドを作ろう|2020/12/06|image,package,walk（まぁまぁ大作に..）
|抽象化|インタフェースを作ろう|2020/12/06|interface(note:interfaceの埋込による設計パターンは深そう)
|エラー処理|エラー処理をしてみよう|2020/12/09|errorインタフェース
|エラー処理|エラー処理をまとめる|2020/12/09|rune,エラー処理
|テストとテスタビリティ|基本構文|2020/11/23|*testing.T
|テストとテスタビリティ|サブテストとテーブル駆動テスト|2020/11/23|サブテスト,Helper関数 [参考](https://qiita.com/nirasan/items/b357f0ad9172ab9fa19b)
|テストとテスタビリティ|単体テストを行おう|2020/11/23|*testing.T[参考](https://engineering.mercari.com/blog/entry/2018-08-08-080000/)