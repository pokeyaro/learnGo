# 字符串格式化输出

## 缺省格式

<table style="font-size: 14px; line-height: 20px;">
<thead>
  <tr>
    <th style="width: auto;">格式化后输出</th>
    <th style="width: auto;">格式符</th>
    <th style="width: auto;">描述</th>
  </tr>
</thead>
<tbody>
  <tr>
    <td>[鸣人 佐助 小樱]</td>
    <td><code>%v</code></td>
    <td>缺省格式</td>
  </tr>
  <tr>
    <td>[]string{"鸣人", "佐助", "小樱"}</td>
    <td><code>%#v</code></td>
    <td>Go 语法打印</td>
  </tr>
  <tr>
    <td>[]string</td>
    <td><code>%T</code></td>
    <td>类型打印</td>
  </tr>
</tbody>
</table>

## 整型

缩进, 进制类型, 正负符号

<table style="font-size: 14px; line-height: 20px;">
<thead>
  <tr>
    <th style="width: auto;">格式化后输出</th>
    <th style="width: auto;">格式符</th>
    <th style="width: auto;">描述</th>
  </tr>
</thead>
<tbody>
  <tr>
    <td>28</td>
    <td><code>%d</code></td>
    <td>十进制</td>
  </tr>
  <tr>
    <td>+28</td>
    <td><code>%+d</code></td>
    <td>必须显示正负符号</td>
  </tr>
  <tr>
    <td>␣␣28</td>
    <td><code>%4d</code></td>
    <td>Pad 空格 (宽度为 4，右对齐)</td>
  </tr>
  <tr>
    <td>28␣␣</td>
    <td><code>%-4d</code></td>
    <td>Pad 空格 (宽度为 4，左对齐)</td>
  </tr>
  <tr>
    <td>11100</td>
    <td><code>%b</code></td>
    <td>二进制</td>
  </tr>
  <tr>
    <td>34</td>
    <td><code>%o</code></td>
    <td>八进制</td>
  </tr>
  <tr>
    <td>1c</td>
    <td><code>%x</code></td>
    <td>十六进制 (小写)</td>
  </tr>
</tbody>
</table>

## 字符

Unicode letter

<table style="font-size: 14px; line-height: 20px;">
<thead>
  <tr>
    <th style="width: auto;">格式化后输出</th>
    <th style="width: auto;">格式符</th>
    <th style="width: auto;">描述</th>
  </tr>
</thead>
<tbody>
  <tr>
    <td>🥷</td>
    <td><code>%c</code></td>
    <td>字符</td>
  </tr>
  <tr>
    <td>'🥷'</td>
    <td><code>%q</code></td>
    <td>显示引号的字符</td>
  </tr>
  <tr>
    <td>U+1F977</td>
    <td><code>%U</code></td>
    <td>Unicode 编码</td>
  </tr>
  <tr>
    <td>U+1F977 '🥷'</td>
    <td><code>%#U</code></td>
    <td>Unicode 编码和显示引号的字符</td>
  </tr>
</tbody>
</table>

## 浮点型

缩进, 精度, 科学计数

<table style="font-size: 14px; line-height: 20px;">
<thead>
  <tr>
    <th style="width: auto;">格式化后输出</th>
    <th style="width: auto;">格式符</th>
    <th style="width: auto;">描述</th>
  </tr>
</thead>
<tbody>
  <tr>
    <td>3.141500e+00</td>
    <td><code>%e</code></td>
    <td>科学计数</td>
  </tr>
  <tr>
    <td>3.141500</td>
    <td><code>%f</code></td>
    <td>十进制小数</td>
  </tr>
</tbody>
</table>

## 字符串 或 []byte

<table style="font-size: 14px; line-height: 20px;">
<thead>
  <tr>
    <th style="width: auto;">格式化后输出</th>
    <th style="width: auto;">格式符</th>
    <th style="width: auto;">描述</th>
  </tr>
</thead>
<tbody>
  <tr>
    <td>Naruto</td>
    <td><code>%s</code></td>
    <td>字符串原样输出</td>
  </tr>
  <tr>
    <td>␣␣␣␣Naruto</td>
    <td><code>%10s</code></td>
    <td>宽度为 10，右对齐</td>
  </tr>
</tbody>
</table>
