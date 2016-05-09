$(document).ready(function() {

	$("#email").on("focusout", function() {
		validateEmail(this);
		if(!$(".errorMessage").is(":visible")) {
			isUserTaken($("#email").val());
		}
	});

	$("#password").on("focusout", function() {
		validatePassword(this);
	});

	$("#retryPassword").on("focusout", function() {
		isPasswordSame(true);
	});
	$("#fname").on("focusout", function() {
		if($(this).val() == "") showError("First Name cannot be blank");
	});
	$("#lname").on("focusout", function() {
		if($(this).val() == "") showError("Last Name cannot be blank");
	});
	$(".submit").on("mousedown", function(e) {
		validateEmail($("#email"));
		if(!$(".errorMessage").is(":visible")) {
			validatePassword($("#password"));
			if(!$(".errorMessage").is(":visible")) {
				isPasswordSame(false);
				if(!$(".errorMessage").is(":visible")) {
					if($("#fname").val() == "") {
						showError("First Name cannot be blank");
					} else {
						if($("#lname").val() == "") {
							showError("Last Name cannot be blank");
						} else {
							removeError();
						}
					}
				}
			}
		}
	});
	$("#resetBtn").on("mouseDown", function(e) {
		e.preventDefault();
		removeError();
		$('signUpForm').reset();
	});
});

function isUserTaken(username) {
	$.ajax({
		url : "/signup/isusertaken",
		data : {
			"username" : username
		},
		success : function(result) {
			if (result.includes('true')) {
				showError("User is already taken !!");
			} else {
				showSuccess("Username available !!");
			}
		}
	});
}