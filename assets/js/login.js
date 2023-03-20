$("#login-button").click(function(event){
    event.preventDefault();

    $.post("/login",
    {
        username: $('#username').val(),
        password: $('#password').val()
    },
    function(data,status) {
        //alert(status) // true/false
        //alert(JSON.stringify(data));
        if (data.success) {
            location.href = "/";
        } else {
            $('#error').text(data.message);
            $("#error").fadeIn();
            $("#error").delay(3000).fadeOut();
        }
    });

    //$('form').fadeOut(500);
    //$('.wrapper').addClass('form-success');
});