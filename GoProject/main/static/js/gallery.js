$(document).ready(function() {
    $("#file").change(function () {
        validateFileFormat(this);
    });
    $(".reset").on("click", function() {
       removeError();
    });
    $(".submit").on("mousedown", function(event) {
        validateFileFormat($("#file"));
        if(!$(".errorMessage").is(":visible")) {
            $("#galleryForm").submit();
        }
    });
    $('.grid').masonry({
        itemSelector: '.grid-item',
        columnWidth: 200
    });
    setTimeout(function() {
        $('.grid').masonry({
            itemSelector: '.grid-item',
            columnWidth: 200
        });
    }, 1500);
});