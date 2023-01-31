# 字符串拼接方式

Go 语言中的三种字符串拼接方式对比

<table style="font-size: 14px; line-height: 20px;">
<thead>
  <tr>
    <th style="width: auto;">方式</th>
    <th style="width: auto;">性能</th>
    <th style="width: auto;">建议</th>
    <th style="width: auto;">理由</th>
  </tr>
</thead>
<tbody>
  <tr>
    <td>Println 加号拼接</td>
    <td>性能一般</td>
    <td>不推荐</td>
    <td>代码难以维护，不便于阅读，可用于简易拼接场景</td>
  </tr>
  <tr>
    <td>Printf 或 Sprintf 格式化</td>
    <td>性能差</td>
    <td>推荐</td>
    <td>十分常用，满足大部分日常需求的处理场景</td>
  </tr>
  <tr>
    <td>Builder 方式</td>
    <td>性能好</td>
    <td>推荐</td>
    <td>不常用，仅对性能要求高的场景下使用</td>
  </tr>
</tbody>
</table>
