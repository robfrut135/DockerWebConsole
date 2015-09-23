<div id="content">
  <h1>Reset password</h1>
  <br>
  {{if .flash.error}}
  <h3>{{.flash.error}}</h3>
  <br>
  {{end}}{{if .flash.notice}}
  <h3>{{.flash.notice}}</h3>
  <br>
  {{end}}
  <form method="POST">

    <div class="form-group">
       {{if .Errors.password}}{{.Errors.password}}{{end}}
       <label for="password">New password (must be at least 6 characters):</label>
       <input name="password" type="password" class="form-control"/>

       {{if .Errors.password2}}{{.Errors.password2}}{{end}}
       <label for="password2">Confirm new password:</label>
       <input name="password2" type="password" class="form-control"/>
    </div>

    <input type="submit" value="Reset password" class="btn btn-default" />

  </form>
</div>
