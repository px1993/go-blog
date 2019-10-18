<html>
<head>
	<title>用户登录</title>
</head>
<style>
</style>
<body>
	<form action="/auth/login" method="post">
		用户名：<input name="username" type="text"></br>
		密  码：<input name="password" type="password"></br>
		<input type="submit">
	</form>
	<p>没有账号？立即<a href="/auth">注册</a></p>
</body>
</html>
