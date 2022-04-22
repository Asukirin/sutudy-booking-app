# TODO

- package 化
- 同一ユーザーのチェック追加
- DB で予約情報を管理
- 画面を用意

# Print について

基本形は`Print`で前後に付属する F とか f とかで挙動が変わる

- `Print`はそのまま出力
- `Fprint`系は書き込み先を指定できる<br>
  ※書き込み先とは
- `Sprint`系はフォーマットした結果を文字列で返す
- `Printf`系はフォーマットの指定が可能<br>
  https://xn--go-hh0g6u.com/pkg/fmt/
- `Println`系はオペランド間に半角スペースが入り、文字列の最後に改行が入る

# Scan について

`fmt.Scan(&変数)`でコンソールからの入力を取得できる<br>
&は変数のアドレスを表している<br>
入力値のキャストが不要<br>
改行またはスペースまでを一つの入力として読み込む点に注意<br>
例：入力 →Hello World!<br>
出力 →Hello<br>
`fmt.Scan(&変数[0], &変数[1], &変数[2],)`<br>
`fmt.Scan(&変数[i])※for文使用`<br>
のように記載することでスペース・改行入りの値を受け取ることはできる<br>

<br>

bufio ライブラリを使用することで１行分取得可能<br>
ただし string 型で取得されるため、キャストが必要な場合がある<br>
例：<br>
`func main() {` <br>
`scanner := bufio.NewScanner(os.Stdin)`<br>
`scanner.Scan()`<br>
`fmt.Println(scanner.Text())`<br>
`}`<br>
ファイルを指定することで、テキストファイルを読み込むことも可能

<br>

速度は１万データを超えるまでどちらも変わらないらしい<br>
１０万データを超えると`fmt.Scan`が明らかに遅くなるらしい

# append について

スライスへ新しい要素を追加する処理<br>
`slice = append(slice, "追加したい値")`

<br>

Arrays は固定長で Slices は可変<br>
Arrays→`var 変数名 [長さ]型`<br>
Slices→`var 変数名 []型`

# make について

スライスの型、長さ、容量を指定する<br>
`make([]型, len, cap)`<br>
メモリのお話は後々

# map について

辞書的なもの？<br>
認識合ってるっぽい

# range について

Slices や、Maps をひとつずつ反復処理するために使う<br>

Slices を range で繰り返す場合、range は反復毎に２つの変数を返す<br>
１つ目の変数は index で、２つ目は index に対応した value <br>

Maps を range で繰り返す場合、反復される順番がランダムになるが、それ以外は Slices と同じ<br>
１つ目の変数は key で、２つ目は key に対応した value

# &,ポインタについて

ポインタとはメモリのアドレス情報のこと<br>
ポインタ変数はメモリ上のアドレスを値として入れられる変数のこと<br>
&を使用することでポインタ型を生成できる<br>
以下を参照
https://qiita.com/Sekky0905/items/447efa04a95e3fec217f

# \_について

宣言するけど使わない・なにもしない変数<br>
配列のインデックスは表示しなくて良い場合などに使用

# strings.Fields について

文字列を空白で分割してスライスを返す

# strconv 　について

型変換に使用

# type,struct について

構造体の定義に使用<br>
異なるタイプの値を一緒に扱うことができる<br>
値へのアクセスは`.フィールド名`

# time.Sleep について

処理をスリープさせる

# go について

並行処理に関する何かだった気がする<br>
調べても出てこないので保留

# sync.WaitGroup について

処理の終了を待つための機能的なもの<br>
WaitGroup の値を作成 →.Add で WaitGroup の値をインクリメント → 処理完了時に.Done でデクリメント →.Wait で完了を待つ(値が 0 になるまで待っている？)

# その他

- :=は変数を定義して値を代入
- パッケージの関数の頭文字は大文字
