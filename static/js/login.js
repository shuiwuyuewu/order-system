$(document).ready(function(){
    $("#login-form").submit(function(e){
        e.preventDefault();
        try {
        $.ajax({
            type: $(this).attr("method"),
            url: $(this).attr("action"),
            data: $(this).serialize(),
        }).done(function(data) {
            console.log(data)
            if (data.Result==2) {
                window.location = data.RedirectPath;
            } else {
                // $("#error-message").
                ShowErrorMessage();
            }
        }).fail(function() {
            alert("submit token failed");
        });
    } catch (e) {
        alert(e)
    }
    });
});

function ShowErrorMessage() {
    if ($("#error-message").length===0) {
        $("h2").after("<div class=\"alert alert-danger\" id=\"error-message\"> <a class=\"close\" data-dismiss=\"alert\" href=\"#\">×</a>令牌非法");
    }
}