<!DOCTYPE html>
<html>
  <head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  </head>
  <body>
    <form method="post" action="/" >
      TODO：<input type="text" name="task">
      <input type="submit" name="Submit" value="登録" />
    </form>
    <form method="post" action="/hoge">
      <ul>
        {{ range $task := .todos }}
        <li>
          <input type="checkbox" name="todo" value="{{ $task.Id }}">{{ $task.Todo }}
          <input type="submit" value="削除">
        </li>
        {{ end }}
      </ul>
    </form>
  </body>
</html>
