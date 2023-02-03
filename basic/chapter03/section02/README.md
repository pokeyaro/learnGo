# 转义符与多行字符串

## 转义符号

<table style="font-size: 14px; line-height: 20px;">
<thead>
  <tr>
    <th style="width: auto;">转义字符</th>
    <th style="width: auto;">说明</th>
  </tr>
</thead>
<tbody>
  <tr>
    <td>\n</td>
    <td>换行(LF) ，将当前位置移到下一行开头</td>
  </tr>
  <tr>
    <td>\r</td>
    <td>回车(CR) ，将当前位置移到本行开头</td>
  </tr>
  <tr>
    <td>\t</td>
    <td>水平制表(HT) （跳到下一个TAB位置）</td>
  </tr>
  <tr>
    <td>\\</td>
    <td>代表一个反斜线字符</td>
  </tr>
  <tr>
    <td>\'</td>
    <td>代表一个单引号（撇号）字符</td>
  </tr>
  <tr>
    <td>\"</td>
    <td>代表一个双引号字符</td>
  </tr>
  <tr>
    <td>\?</td>
    <td>代表一个问号</td>
  </tr>
</tbody>
</table>

## 多行字符串

```go
func main() {
    // 单行字符串使用 ""
    msg1 := "hello ~\r\n\"naruto\""

    // 多行字符串使用 ``
    msg2 := `hello ~
"naruto"`
}
```
