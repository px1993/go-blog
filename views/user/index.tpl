<html>
<head>
	<title>个人中心</title>
	<style>
		table td{border:1px solid #000} 
		/* css注释：只对table标签设置红色边框样式 */ 
	</style>
</head>
<body>
	<h1>{{ .username }}的个人首页</h1>
	<p><a href="/auth/logout">退出</a></p>
	<table>
		<tr>
			<th width="105">用户Id</th>
			<th width="105">用户名</th>
			<th width="105">用户密码</th>
		<tr>
		<tr>
			<td>{{ .userId }}</td>
			<td>{{ .username }}</td>
			<td>{{ .password }}</td>
		</tr>
	</table>
</body>
</html>
