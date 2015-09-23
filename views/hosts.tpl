		<!---start-content---->

				  <ul id="commands">
					<li>
						<div class="main-searchbar">
							<label>Run Docker commands in: localhost </label>
							<input id="command" type="text" placeholder="Your docker command"/>
							<input type="submit" onClick="Service.command($('#command').val()); $('#command').val('Running your command on host...')" value="" />
						</div>
					</li>
				  </ul>
			      <ul id="tiles">
			        {{ range $key, $container := .Containers }}
						<!-- These are our grid blocks -->
				        <li>
				        	<div class="post-info">
				        		<div class="post-basic-info">
											<img title="Web Console" src="/static/images/docker.png" width="200" height="200" onClick="Service.activeGotty({{ $container.Id }})">
					        		<h3><a href="#" title="Web Console" onClick="Service.activeGotty({{ $container.Id }})">{{ $key }}</a></h3>
					        		<span><a href="#"><label> </label>{{ $container.Image }}</a></span>
					        		<p>{{ $container.Names }}</p>
											<p class="container-status">
												<span>{{ $container.Status }} - <a title="Show container logs" href="#" onClick="Service.logs({{ $container.Id }})">[ Log real-time ]</a></span>
											</p>
				        		</div>
				        		<div class="post-info-rate-share">
				        			<div title="Inspect container" class="post-share" onClick="Service.command('inspect ' + {{ $container.Id }})">
				        				<span></span>
				        			</div>
				        			<div class="clear">
											</div>
				        		</div>
				        	</div>
				        </li>
			       	{{ end }}
			      </ul>

		<!---//End-content---->
