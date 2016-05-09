$(document).ready(function() {
    $("#username").on("focusout", function() {
        validateEmail(this);
    });
    $("#password").on("focusout", function() {
        validatePassword(this);
    });
    $(".submit").on("mousedown", function(event) {
        validateEmail($("#username"));
        if(!$(".errorMessage").is(":visible")) {
            validatePassword("#password");
        }
    });
});