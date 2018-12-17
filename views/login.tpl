<!doctype html>
<html lang="ru">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
    {{ if .Error }}
            <p>{{.Error}}</p>
    {{ end }}
    <form action="/login" method="POST" enctype="multipart/form-data">
        <input type="text" name="username_log" placeholder="Username" required>
        <input type="password" name="password" placeholder="Password" required>
        <input type="submit" value="Login">
    </form>
    <form action="/register" method="POST" enctype="multipart/form-data">
        <input type="text" name="username_reg" placeholder="Username" required>
        <input type="password" name="password" placeholder="Password" required>
        <input type="password" name="repassword" placeholder="RepeatPassword" required>
        <input type="submit" value="Register">
    </form>
</body>
</html>