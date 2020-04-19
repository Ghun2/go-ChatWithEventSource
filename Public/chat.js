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
        return false;
    });

    const addMessage = (data) => {
        let text = "";
        if (!isBlank(data.name)) {
            text = '<strong>' + data.name + ':</strong>';
        }
        text += data.msg;
        $chatlog.append('<div><span>'+ text +'</span></div>');
    };

    const es = new EventSource('/stream');
    es.onopen = function(e) {
        $.post('users/', {
            name: username
        });
    };
    es.onmessage = function (e) {
        const msg = JSON.parse(e.data);
        addMessage(msg);
    };

    window.onbeforeunload = function () {
        $.ajax({
            url: "/users?username=" + username,
            type: "DELETE"
        });
        es.close();
    };
})