<div id="content">
  <h1>Add</h1>
  <br>
  {{if .flash.error}}
  <h3>{{.flash.error}}</h3>
  <br>
  {{end}}
  {{if .flash.notice}}
  <h3>{{.flash.notice}}</h3>
  <br>
  {{end}}
  <form method="POST">
    <div class="form-group">
      <label>Id:</label>
      <input name="-" value="{{.User.Id}}" size="10" readonly class="form-control"/>

      {{if .Errors.First}}{{.Errors.First}}{{end}}
      <label for="first">First: </label>
      <input name="first" value="{{.User.First}}" size="30" class="form-control"/>

      <label for="last">Last:</label>
      <input name="last" value="{{.User.Last}}" size="30" class="form-control"/>

      {{if .Errors.Email}}{{.Errors.Email}}{{end}}
      <label for="email">Email:</label>
      <input name="email" value="{{.User.Email}}" size="30" class="form-control" />

      {{if .Errors.Password}}{{.Errors.Password}}{{end}}
      <label for="password">Password:</label>
      <input name="password" value="{{.User.Password}}" size="30" class="form-control"/>

      <label for="reg_key">Reg key</label>
      <input name="reg_key" value="{{.User.Reg_key}}" size="50" readonly class="form-control" />

      <label for="reg_date">Reg date</label>
      <input name="reg_date" value="{{.User.Reg_date}}" size="50" readonly class="form-control"/>

      <label for="reset_key">Reset key</label>
      <input name="reset_key" value="{{.User.Reset_key}}" size="50" readonly class="form-control" />

    </div>

    <input type="submit" value="Add" class="btn btn-default"  /></td>

  </form>
  <h4><a href="http://{{.domainname}}/appadmin/index/{{.parms}}">Return to Index</a></h4>
</div>
