
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
          {{if .Errors.email}}{{.Errors.email}}{{end}}
  				<input id="email" type="text" placeholder="Your email">
          {{if .Errors.password}}{{.Errors.password}}{{end}}
  				<input id="password" type="password" placeholder="Your password">

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
