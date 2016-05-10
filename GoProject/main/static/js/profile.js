$(document).ready(function() {
    removeError();
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
        validatePassword($("#password"));
        if(!$(".errorMessage").is(":visible")) {
            isPasswordSame(false);
            if (!$(".errorMessage").is(":visible")) {
                if ($("#fname").val() == "") {
                    showError("First Name cannot be blank");
                } else {
                    if ($("#lname").val() == "") {
                        showError("Last Name cannot be blank");
                    } else {
                        removeError();
                    }
                }
            }
        }
    });
    $("#resetBtn").on("click", function() {
        removeError();
        $('signUpForm').reset()
    });
});
