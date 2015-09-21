 $("#login-button").click(function(event){

	$("form").submit();

	event.preventDefault();

	$('form').fadeOut(500);
	$('.wrapper').addClass('form-success');
});
