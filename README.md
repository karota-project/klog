klog
============

## Feature
- benchmark for roomba2d2
- logging cpu & mem usage
- use "log" "log/syslog" pkg 

## ToDo

## Usage
- 定義

<pre>
func Stdout(_func string) (err error) 
func Syslog(p Priority, facility string) (err error) 
func WriteFile(_func string, outfile string) (err error) 
</pre>

- 使い方

<pre>
err := klog.Stdout("main")
if err != nil {
	fmt.Println(err)
}

err = klog.WriteFile("main", "sample.log")
if err != nil {
	fmt.Println(err)
}

err = klog.Syslog(klog.LOG_NOTICE, "main")
if err != nil {
	fmt.Println(err)
}
</pre>

- 標準出力

<pre>
[2014-06-07 10:58:20.300722544 +0900 JST] /klog-master/klog_example.go(line14) {"func" : "main" ,"mem_used" : 379220, "mem_free" : 480440, "cpu_used" : 2}
</pre>

