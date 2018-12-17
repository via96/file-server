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
    {{if .CurrentUser}}
        <form action="/logout" method="POST" enctype="multipart/form-data">
            <label>{{.CurrentUser.Login}}</label>
            <input type="submit" value="Logout">
        </form>
    {{end}}

    {{ if .Error }}
        <p>{{.Error}}</p>
    {{ end }}
    <form action="/upload" method="POST" enctype="multipart/form-data">
        <input type="file" name="file_loader" multiple>
        <input type="submit">
    </form>
    <table>
    {{ if .Files}}
        <thead>
        <tr>
            <td>File name</td>
            <td>Upload time</td>
            <td></td>
            <td></td>
        </tr>
        </thead>
        <tbody>
        {{range $key, $val := .Files}}
        <tr>

            <td>{{$val.UserFileName}}</td>
            <td>{{$val.UploadTime}}</td>
            <td><a href="/download/{{$val.Id}}">Download</a></td>

            <td><a href="/remove/{{$val.Id}}">Remove</a></td>
        </tr>
        {{end}}
        </tbody>
    {{end}}

    </table>
</body>
</html>