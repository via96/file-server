<!doctype html>
<html lang="ru">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>File Server</title>
    <!-- Bootstrap -->
    <link href="/static/css/bootstrap.min.css" rel="stylesheet">
    <link href="/static/css/bootstrap-theme.min.css" rel="stylesheet">
    <link href="/static/css/bootstrap-responsive.css" rel="stylesheet">
    <script src="/static/js/jquery.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.12.9/umd/popper.min.js"></script>
    <script src="/static/js/bootstrap.min.js"></script>
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