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
    <div class="container-fluid">
        <div class="row" style="background-color: #169d5d; margin-bottom: 20px;">
            <a href="/"><img src="/static/img/logo_main.png" height="60" alt="Тут должна быть картинка"></a>
        </div>

        <div class="row">
            <div style="border: #20c997 3px solid; border-radius: 10px" class="col-3 offset-4">
                <form action="/login" method="POST" enctype="multipart/form-data">
                    {{ if .LoginError }}
                        <br>
                        <div class="alert alert-danger">{{.LoginError}}</div>
                    {{ else }}
                        <br><br>
                    {{ end }}
                    <div class="form-group">
                        <label for="name_log">Имя пользователя</label>
                        <input id="name_log" class="form-control" type="text" name="username_log" placeholder="Username" required>
                    </div>
                    <div class="form-group">
                        <label for="pwd_log">Пароль</label>
                        <input id="pwd_log" class="form-control" type="password" name="password" placeholder="Password" required>
                    </div>
                    <div align="right"><input class="btn btn-success" type="submit" value="Login"></div>
                </form>
            </div>
            <div style="border: #20c997 3px solid; border-radius: 10px" class="col-2 offset-1">
                <form action="/register" method="POST" enctype="multipart/form-data">
                    {{ if .RegisterError }}
                        <div class="alert alert-danger">{{.RegisterError}}</div>
                    {{ else }}
                        <br>
                    {{ end }}
                    <div class="form-group">
                        <label for="name_reg">Имя пользователя</label>
                        <input id="name_reg" class="form-control" type="text" name="username_reg" placeholder="Username" required>
                    </div>
                    <div class="form-group">
                        <label for="pwd_reg">Пароль</label>
                        <input id="pwd_reg" class="form-control" type="password" name="password" placeholder="Password" required>
                    </div>
                    <div class="form-group">
                        <label for="repwd_reg">Подтверждение пароля</label>
                        <input id="repwd_reg" class="form-control" type="password" name="repassword" placeholder="RepeatPassword" required>
                    </div>
                    <div align="right"><input class="btn btn-success" type="submit" value="Register"></div><br>
                </form>
            </div>
        </div>
    </div>

</body>
</html>