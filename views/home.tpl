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
    {{if .currentUser}}
        <label>{{.currentUserName}}</label><br>
        <input type="button" value="Logout">
    {{else}}
        <form action="/login" method="POST" enctype="multipart/form-data">
            <input type="text" name="username" placeholder="Username" required>
            <input type="password" name="password" placeholder="Password" required>
            <input type="submit" value="Login">
        </form>
    {{end}}

    <form action="/" method="POST" enctype="multipart/form-data">
        <input type="file" name="file_loader" multiple>
        <input type="submit">
    </form>

</body>
</html>