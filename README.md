# kingpin お試し

kingpinについては、以下を参照のこと。

http://qiita.com/kumatch/items/258d7984c0270f6dd73a

## 基本的な使い方

以下のようにvarを作成することで、Flag(--から始まるオプション)やArg(Flag以外)を受け取ることができます。
また、デフォルトで`--help`は受け取ることが可能ですが`-h`は受とれないため、`kingpin.CommandLine.HelpFlag.Short('h')`で設定しています。

デフォルト値は、型毎の0値が設定されますが`Default()`にて上書き可能です。
その際、`Default()`の引数はstringのため、`bool`値を上書きする場合は"true"とかで設定する必要があります。

デフォルト値は、`OverrideDefaultFromEnvar()`によって上書き可能です。

```golang
package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	verbose = kingpin.Flag("verbose", "Set verbose mode").Short('v').OverrideDefaultFromEnvar("KINGPIN_SAMPLE_VERBOSE").Bool()
	count   = kingpin.Flag("count", "counter").Default("10").Short('c').Int()
	name    = kingpin.Arg("name", "Input name").Required().String()
)

func main() {
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	fmt.Printf("verbose mode: %v\n", *verbose)
	fmt.Printf("count       : %d\n", *count)
	fmt.Printf("name        : %s\n", *name)
}
```
