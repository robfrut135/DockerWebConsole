    <div class="wrapper">

  		<div class="container">
  			<h1>Docker Web Console</h1>

        <br>
        {{if .flash.error}}
        <h3>{{.flash.error}}</h3>
        <br>
        {{end}}

  			<form id="user-data" class="form" method="post">
          <div class="form-group">
            {{if .Errors.email}}{{.Errors.email}}{{end}}
            <label for="email">Email address</label>
    				<input id="email" name="email" type="text" class="form-control" placeholder="Your email">

            {{if .Errors.password}}{{.Errors.password}}{{end}}
            <label for="password">Password</label>
            <input id="password" name="password" class="form-control" type="password" placeholder="Your password">
          </div>
  				<button type="submit" id="login-button" class="btn btn-default">Login</button>

          <div>
            <br>
            <a title="Register" href="http://{{.domainname}}/user/register">Register</a> || <a href="http://{{.domainname}}/user/forgot">Forgot password?</a>
          </div>

  			</form>
  		</div>

  		<ul class="bg-bubbles">
  			<li></li>
  			<li></li>
  			<li></li>
  			<li></li>
  			<li></li>
  			<li></li>
  			<li></li>
  			<li></li>
  			<li></li>
  			<li></li>
  		</ul>
	  </div>

    <script src="/static/js/login.js"></script>
