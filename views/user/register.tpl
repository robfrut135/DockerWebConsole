<div id="content">
    <h1>Please register</h1>
    <br>
    {{if .flash.error}}
    <h3>{{.flash.error}}</h3>
    <br>
    {{end}}
    <form method="POST">
      <div class="form-group">
        {{if .Errors.First}}{{.Errors.First}}{{end}}
        <label for="first">First name: </label>
        <input name="first" type="text" value="{{.User.First}}" autofocus class="form-control"/>

        <label for="last">Last name:</label>
        <input name="last" type="text" value="{{.User.Last}}"class="form-control" />

        {{if .Errors.Email}}{{.Errors.Email}}{{end}}
        <label for="email">Email address:</label>
        <input name="email" type="text" value="{{.User.Email}}" class="form-control"/>

         {{if .Errors.Password}}{{.Errors.Password}}{{end}}
        <label for="password">Password (must be at least 6 characters):</label>
        <input name="password" type="password" class="form-control"/>

        {{if .Errors.Confirm}}{{.Errors.Confirm}}{{end}}
        <label for="password2">Confirm password: </label>
        <input name="password2" type="password" class="form-control"/>
      </div>

      <input type="submit" value="Register" class="btn btn-default"/>

    </form>
</div>
