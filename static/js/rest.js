var Service = {

	activeGotty		: function (idContainer){
						var url = "/hosts/console/intro?id=" + idContainer

						$.ajax({
							dataType: "text",
							type: "GET",
							url: url,
							beforeSend: function(){
								$("#demo01").trigger("click");
							}
						})
						.always(function(){
							$("#animatedModal .close-animatedModal").trigger("click");
						});

						setTimeout(	"window.open('http://localhost:9999','winname','directories=0,titlebar=0,toolbar=0,location=0,status=0,menubar=0,scrollbars=no,resizable=no')",
									2000
									);
					},
	logs		: function (idContainer){
						var url = "/hosts/console/logs?id=" + idContainer

						$.ajax({
							dataType: "text",
							type: "GET",
							url: url,
							beforeSend: function(){
								$("#demo01").trigger("click");
							}
						})
						.always(function(){
							$("#animatedModal .close-animatedModal").trigger("click");
						});

						setTimeout(	"window.open('http://localhost:8888','winname','directories=0,titlebar=0,toolbar=0,location=0,status=0,menubar=0,scrollbars=no,resizable=no,width=800,height=400')",
									2000
									);
					},
	command		: function (cmd){
						var url = "/hosts/console/cmd?command=" + cmd

						$.ajax({
							dataType: "text",
							type: "GET",
							url: url,
						})
						setTimeout(	"window.open('http://localhost:7777','winname','directories=0,titlebar=0,toolbar=0,location=0,status=0,menubar=0,scrollbars=no,resizable=no,width=800,height=400')",
									2000
									);
					},
};

$(document).ready(function(){

	$("#demo01").animatedModal({
        modalTarget:'animatedModal',
        animatedIn:'zoomIn',
        animatedOut:'bounceOutDown',
        color:'grey',
        // Callbacks
        beforeOpen: function() {
            console.log("The animation was called");
        },
        afterOpen: function() {
            console.log("The animation is completed");
        },
        beforeClose: function() {
            console.log("The animation was called");
        },
        afterClose: function() {
            console.log("The animation is completed");
        }
    });

});
