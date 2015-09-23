<div id="content">
    <h1>Remove User Account</h1>
    <br>
    {{if .flash.error}}
    <h3>{{.flash.error}}</h3>
    <br>
    {{end}}
    <p><font size="3">Caution: all related transactions and data will also be removed. Are you sure?</font></p>
    <form method="POST">
        <div class="form-group">
          {{if .Errors.current}}{{.Errors.current}}{{end}}
          <label for="current">Current password:</label>
          <td><input name="current" type="password" class="form-control"/></td>
        </div>

        <input type="submit" value="Remove" class="btn btn-default" />

        <a href="http://localhost:8080/home">Cancel</a>

    </form>
</div>
