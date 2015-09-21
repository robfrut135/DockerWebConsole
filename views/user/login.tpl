
    <!-- ToDo: <link rel="stylesheet" href="/static/css/login.css"> -->

    <div class="wrapper">

  		<div class="container">
  			<h1>Docker Web Console</h1>

        <br>
        {{if .flash.error}}
        <h3>{{.flash.error}}</h3>
        <br>
        {{end}}

  			<form id="user-data" class="form" method="post">
          {{if .Errors.password}}{{.Errors.password}}{{end}}
  				<input id="username" type="text" placeholder="Username">
          {{if .Errors.password}}{{.Errors.password}}{{end}}
  				<input id="password" type="password" placeholder="Password">

  				<button type="submit" id="login-button">Login</button>

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
