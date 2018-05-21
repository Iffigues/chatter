<html>
  <head>
    <title>Upload file</title>
  </head>
  <body>
    <form enctype="multipart/form-data" action="http://gopiko.fr/up/t" method="post">
      <input type="email" name="Prenom"/>
      <input type="email" name="email" />

      <input type="email" name="email" />
      <input type="file" name="uploadfile" />
      <input type="hidden" name="token" value="{{.}}"/>
      <input type="submit" value="upload" />
    </form>
  </body>
</html>
