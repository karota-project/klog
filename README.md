klog
============

## Feature
- ルンバプロジェクトのベンチマーク用 (karota-roomba-hack benchmark)
- CPU使用率/ メモリ使用率をロギング (logging cpu & mem usage)
- logパッケージのラッパ (wrapper package)
- Linux向け (for linux)

## ToDo
- ファイル出力で追記モードにする
- ファイル出力時に時刻を入れる

## Usage
- 定義

<pre>
func Printlog(_func string) (result bool, err error) 
</pre>

- 使い方

<pre>
s , err := klog.Printlog("function-name")
if err != nil {
	fmt.Println(s, err)
}

s , err = klog.Printfile("function-name", "output-file-name")
if err != nil {
	fmt.Println(s, err)
}
</pre>

- 標準出力

<pre>
2014/05/08 04:05:22 klog_example.go  , function-name ,mem-used :  557212 kB ,mem-free :  122896 kB ,cpu-used :  1 ％
</pre>

- ファイル出力

<pre>
time : 2014-05-08 04:09:26.045481644 +0900 JST,file :klog_example.go,func : function-name ,mem-used : 564040kB ,mem-free : 107936kB ,cpu-used : 1％
</pre>
