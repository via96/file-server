<!doctype html>
<html lang="ru">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>BeeFileServer</title>
    <link rel="shortcut icon" type="image/png" href="/static/img/favicon.ico"/>
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
            <div class="col-5">
                <a href="/"><img src="/static/img/logo_main.png" height="60" alt="Тут должна быть картинка"></a>
            </div>
            <div class="col-3 offset-4">
                {{if .CurrentUser}}
                    <form action="/logout" method="POST" enctype="multipart/form-data">
                        <label>{{.CurrentUser.Login}}</label>
                        <input id="logout_btn" class="btn btn-dark" type="submit" value="Выход">
                    </form>
                {{end}}
            </div>
        </div>

        <div class="row" style="margin-bottom: 20px">
            <div class="col-6 offset-3">
                <form action="/upload" method="POST" enctype="multipart/form-data">
                    {{ if .Error }}
                        <p>{{.Error}}</p>
                    {{ end }}
                    <div align="center">
                        <input type="file" name="file_loader" multiple>
                        <input class="btn btn-info" type="submit" value="Загрузить">
                    </div>
                </form>
            </div>
        </div>

        <div class="row">
            <div class="col-8 offset-2">
                {{ if .Files}}
                    <table class="table">

                        <thead class="thead-dark">
                        <tr>
                            <th scope="col">Имя файла</th>
                            <th scope="col">Время загрузки</th>
                            <th scope="col"></th>
                            <th scope="col"></th>
                        </tr>
                        </thead>
                        <tbody>
                        {{range $key, $val := .Files}}
                        <tr>

                            <td>{{$val.UserFileName}}</td>
                            <td>{{$val.UploadTime}}</td>
                            <td><a href="/download/{{$val.Id}}">Скачать</a></td>

                            <td><a href="/remove/{{$val.Id}}">Удалить</a></td>
                        </tr>
                        {{end}}
                        </tbody>


                    </table>
                {{ else }}
                    <div class="alert alert-dark" role="alert">
                        <div align="center">Список файлов пуст</div>
                    </div>
                {{end}}



            </div>
        </div>

    </div>
</body>
</html>