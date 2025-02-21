<div align="center">

# Ollama Scanner (扫描器)

</div>

一个多线程的Ollama扫描器工具
使用ants库实现的简单扫描器

## Build (构建)

```text
    go build .
```

## How to use?

```text
    ollamaScan
        -t Num Threads (Default:50)
        -l IP List File(Single IP Line) (Default:ips.txt)
        -o Output File (Default: out.txt)
        -json Save as Json Format (Default: False)
```
```text
    ollamaScan
        -t 线程数 (默认:50)
        -l IP列表文件(一行一个) (默认:ips.txt)
        -o 输出文件 (默认: out.txt)
        -json 保存为Json (默认: False)
```