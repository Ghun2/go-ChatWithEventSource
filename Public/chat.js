$(function () {
    if (!window.EventSource) {
        alert("No EventSource!")
        return
    }

    const $chatlog = $('#chat-log')
    const $chatmsg = $('#chat-msg')

    const isBlank = (str) => {
        return str == null || str.trim() === "";
    };
    let username;
    while (isBlank(username)) {
        username = prompt("What's your name?")
        if (!isBlank(username)) {
            $('#user-name').html('<b>' + username + '</b>')
        }
    }

    $('#input-form').on('submit', () => {
        $.post('/messages', {
            msg: $chatmsg.val(),
            name: username
        });
        $chatmsg.val("");
        $chatmsg.focus();
    })
})