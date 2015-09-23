<div id="content">
<h1>Reset password</h1>
<br>
{{if .flash.error}}
<h3>{{.flash.error}}</h3>
<br>
{{end}}

<form method="POST">
  <div class="form-group">
      {{if .Errors.email}}{{.Errors.email}}{{end}}
      <label for="email">Email address:</label>
      <input name="email" type="text" autofocus class="form-control"/>
  </div>

  <input type="submit" value="Request reset" class="btn btn-default" /></td>

</form>
</div>
