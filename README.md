klog
============

## Feature
- ルンバプロジェクトのベンチマーク用 (karota-roomba-hack benchmark)
- CPU使用率/ メモリ使用率をロギング (logging cpu & mem usage)
- logパッケージのラッパ (wrapper package)
- Linux向け (for linux)

## ToDo
- ファイル出力

## Usage
- 定義

<pre>
func Printlog(fname string) int
</pre>

- 使い方

<pre>
package main

import (
  "./klog"
)

func main () {
	klog.Printlog("function-name")
}
</pre>

- 標準出力

<pre>
2014/04/28 07:00:34 [ function-name ] ,mem-used :  524004 ,mem-free :  105696 ,cpu-used :  1
</pre>
