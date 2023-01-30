# 运算符和表达式

## 算数运算符

<table style="font-size: 14px; line-height: 20px;">
<thead>
  <tr>
    <th style="width: auto;">符号</th>
    <th style="width: auto;">名称</th>
  </tr>
</thead>
<tbody>
  <tr>
    <td>+</td>
    <td>加</td>
  </tr>
  <tr>
    <td>-</td>
    <td>减</td>
  </tr>
  <tr>
    <td>*</td>
    <td>乘</td>
  </tr>
  <tr>
    <td>/</td>
    <td>除</td>
  </tr>
  <tr>
    <td>%</td>
    <td>余</td>
  </tr>
  <tr>
    <td>++</td>
    <td>自增</td>
  </tr>
  <tr>
    <td>--</td>
    <td>自减</td>
  </tr>
</tbody>
</table>


## 关系运算符

<table style="font-size: 14px; line-height: 20px;">
<thead>
  <tr>
    <th style="width: auto;">符号</th>
    <th style="width: auto;">名称</th>
  </tr>
</thead>
<tbody>
  <tr>
    <td>==</td>
    <td>等于</td>
  </tr>
  <tr>
    <td>!=</td>
    <td>不等于</td>
  </tr>
  <tr>
    <td>></td>
    <td>大于</td>
  </tr>
  <tr>
    <td><</td>
    <td>小于</td>
  </tr>
  <tr>
    <td>>=</td>
    <td>大于等于</td>
  </tr>
  <tr>
    <td><=</td>
    <td>小于等于</td>
  </tr>
</tbody>
</table>

## 逻辑运算符

<table style="font-size: 14px; line-height: 20px;">
<thead>
  <tr>
    <th style="width: auto;">符号</th>
    <th style="width: auto;">名称</th>
  </tr>
</thead>
<tbody>
  <tr>
    <td>&&</td>
    <td>逻辑与</td>
  </tr>
  <tr>
    <td>||</td>
    <td>逻辑或</td>
  </tr>
  <tr>
    <td>!</td>
    <td>逻辑非</td>
  </tr>
</tbody>
</table>

## 位运算符

位运算符对整数在内存中的二进制位进行操作

| a   | b   | a & b | a｜b | a ^ b |
|-----|-----|-------|-----|-------|
| 0   | 0   | 0     | 0   | 0     |
| 0   | 1   | 0     | 1   | 1     |
| 1   | 1   | 1     | 1   | 0     |
| 1   | 0   | 0     | 1   | 1     |

## 赋值运算符

<table style="font-size: 14px; line-height: 20px;">
<thead>
  <tr>
    <th style="width: auto;">符号</th>
    <th style="width: auto;">名称</th>
  </tr>
</thead>
<tbody>
  <tr>
    <td>=</td>
    <td>将表达式结果赋给左值</td>
  </tr>
  <tr>
    <td>+=</td>
    <td>相加后再赋值</td>
  </tr>
  <tr>
    <td>-=</td>
    <td>相减后再赋值</td>
  </tr>
  <tr>
    <td>...</td>
    <td>...</td>
  </tr>
</tbody>
</table>

## 其他

- 返回变量存储地址：`&`
- 指针变量：`*`

## 运算符优先级

可以通过括号来提升某个表达式的整体运算优先级

<table style="font-size: 14px; line-height: 20px;">
<thead>
  <tr>
    <th style="width: auto;">优先级</th>
    <th style="width: auto;">分类</th>
    <th style="width: auto;">运算符</th>
    <th style="width: auto;">结合性</th>
  </tr>
</thead>
<tbody>
  <tr>
    <td>01</td>
    <td>逗号运算符</td>
    <td>,</td>
    <td>从左到右</td>
  </tr>
  <tr>
    <td>02</td>
    <td>赋值运算符</td>
    <td>=、+=、-=、*=、/=、%=、>>=、<<=、&=、|=、^=</td>
    <td>从右到左</td>
  </tr>
  <tr>
    <td>03</td>
    <td>逻辑或</td>
    <td>||</td>
    <td>从左到右</td>
  </tr>
  <tr>
    <td>04</td>
    <td>逻辑与</td>
    <td>&&</td>
    <td>从左到右</td>
  </tr>
  <tr>
    <td>05</td>
    <td>按位或</td>
    <td>|</td>
    <td>从左到右</td>
  </tr>
  <tr>
    <td>06</td>
    <td>按位异或</td>
    <td>^</td>
    <td>从左到右</td>
  </tr>
  <tr>
    <td>07</td>
    <td>按位与</td>
    <td>&</td>
    <td>从左到右</td>
  </tr>
  <tr>
    <td>08</td>
    <td>相等/不等</td>
    <td>==、!=</td>
    <td>从左到右</td>
  </tr>
  <tr>
    <td>09</td>
    <td>关系运算符</td>
    <td><、<=、>、>=</td>
    <td>从左到右</td>
  </tr>
  <tr>
    <td>10</td>
    <td>位移运算符</td>
    <td><<、>></td>
    <td>从左到右</td>
  </tr>
  <tr>
    <td>11</td>
    <td>加法/减法</td>
    <td>+、-</td>
    <td>从左到右</td>
  </tr>
  <tr>
    <td>12</td>
    <td>乘法/除法/取余</td>
    <td>*（乘号）、/、%</td>
    <td>从左到右</td>
  </tr>
  <tr>
    <td>13</td>
    <td>单目运算符</td>
    <td>!、*（指针）、& 、++、--、+（正号）、-（负号）</td>
    <td>从右到左</td>
  </tr>
  <tr>
    <td>14</td>
    <td>后缀运算符</td>
    <td>( )、[ ]、-></td>
    <td>从左到右</td>
  </tr>
</tbody>
</table>
