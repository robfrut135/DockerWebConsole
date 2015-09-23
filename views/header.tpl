<!--
Author: W3layouts
Author URL: http://w3layouts.com
License: Creative Commons Attribution 3.0 Unported
License URL: http://creativecommons.org/licenses/by/3.0/
-->
<!DOCTYPE HTML>
<html>
	<head>
		<title>Docker Web Console | Home :: w3layouts</title>
		<link href="/static/css/style.css" rel='stylesheet' type='text/css' />
		<link rel="stylesheet" href="//cdnjs.cloudflare.com/ajax/libs/animate.css/3.2.0/animate.min.css">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<link rel="shortcut icon" type="image/x-icon" href="/static/images/docker.png" />
		<script type="application/x-javascript"> addEventListener("load", function() { setTimeout(hideURLbar, 0); }, false); function hideURLbar(){ window.scrollTo(0,1); } </script>
		</script>
		<!----webfonts---->
		<link href='http://fonts.googleapis.com/css?family=Open+Sans:400,300,600,700,800' rel='stylesheet' type='text/css'>
		<!----//webfonts---->
		<!-- Global CSS for the page and tiles -->
  		<link rel="stylesheet" href="/static/css/main.css">
  		<!-- //Global CSS for the page and tiles -->
		<!---start-click-drop-down-menu----->
		<script src="/static/js/jquery.min.js"></script>
		<!---start animatedModal -->
		<script src="/static/js/animatedModal.min.js"></script>
		<!---start js service -->
		<script src="/static/js/rest.js"></script>
        <!----start-dropdown--->
         <script type="text/javascript">
			var $ = jQuery.noConflict();
				$(function() {
					$('#activator').click(function(){
						$('#box').animate({'top':'0px'},500);
					});
					$('#boxclose').click(function(){
					$('#box').animate({'top':'-700px'},500);
					});
				});
				$(document).ready(function(){
				//Hide (Collapse) the toggle containers on load
				$(".toggle_container").hide();
				//Switch the "Open" and "Close" state per click then slide up/down (depending on open/close state)
				$(".trigger").click(function(){
					$(this).toggleClass("active").next().slideToggle("slow");
						return false; //Prevent the browser jump to the link anchor
				});

			});
		</script>
        <!----//End-dropdown--->
		<!---//End-click-drop-down-menu----->
	</head>
	<body>

				<a id="demo01" href="#animatedModal"></a>
			    <div id="animatedModal">
			        <div class="close-animatedModal">
			        </div>
			        <div class="modal-content">
			            <center>Loading administration console...</center>
			        </div>
			    </div>
				<!---start-wrap---->
					<!---start-header---->
					<div class="header">
						<div class="wrap">
						<div class="logo">
							<a href="/home"><img src="/static/images/docker-logo.png" title="Docker Web Console" width="154px" height="52px" /></a>
						</div>
						<div class="nav-icon">
							 <a href="#" class="right_bt" id="activator"><span></span> </a>
						</div>
						 <div class="box" id="box">
							 <div class="box_content">
								<div class="box_content_center">
								 	<div class="form_content">
										<div class="menu_box_list">
											<ul>
												<li><a href="/home"><span>home</span></a></li>
												<li><a href="/hosts"><span>Hosts</span></a></li>
												<li><a href="/notice"><span>Blog</span></a></li>
												<li><a href="/about"><span>About</span></a></li>
												<div class="clear"> </div>
											</ul>
										</div>
										<a class="boxclose" id="boxclose"> <span> </span></a>
									</div>
								</div>
							</div>
						</div>
						<div class="top-searchbar">
							<form>
								<input id="pattern" name="pattern" type="text" /><input type="submit" value="" />
							</form>
						</div>
						<div class="userinfo">
							<div class="user">
									<ul>
										{{if .InSession}}
												<li><img src="/static/images/user-pic.png" title="user-name" /><span>{{.First}}</span></li>
												<li>[<a href="http://{{.domainname}}/user/logout">Logout</a>|<a href="http://{{.domainname}}/user/profile">Profile</a>]</li>
										{{else}}
												[<a href="http://{{.domainname}}/user/login/hosts">Login</a>]
										{{end}}
									</ul>
							</div>
						</div>
						<div class="clear"> </div>
					</div>
				</div>
				<!---//End-header---->

				<div class="content">
				  <div class="wrap">
				    <div id="main" role="main">
