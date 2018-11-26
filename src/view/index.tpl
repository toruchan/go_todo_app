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
    <ul>
      {{ range $task := .todos }}
      <li>
        <input type="checkbox" name="todo" value="{{ $task }}">{{ $task }}
      </li>
      {{ end }}
    </ul>
  </body>
</html>
