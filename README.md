klog
============

## Feature
- ルンバプロジェクトのベンチマーク用 (karota-roomba-hack benchmark)
- CPU使用率/ メモリ使用率をロギング (logging cpu & mem usage)
- logパッケージのラッパ (wrapper package)
- Linux向け (for linux)

## ToDo

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
[2014-06-07 10:58:20.300722544 +0900 JST] /klog-master/klog_example.go(line14) {"func" : "main" ,"mem_used" : 379220, "mem_free" : 480440, "cpu_used" : 2}
</pre>

