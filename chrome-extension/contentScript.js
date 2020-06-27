// construct an HTTP request
let xhr = new XMLHttpRequest();
xhr.open('POST', 'http://127.0.0.1:60606/api/v1/visit/' + window.location.href, true);

xhr.send(document.body.innerHTML);
