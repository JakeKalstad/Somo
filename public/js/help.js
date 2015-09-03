function log(msg, sev) {
	'use strict';	
	if (!sev) {
		console.log(sev);
	} else {
		//client side error log to server
	}
}

function displayErrors(errors) {
	'use strict';
	for (var idx in errors) {
		var error = errors[idx];
		var label = $(".err_" + error.type + ":visible");
		label.parent().tooltip("destroy")
		label.parent().data("title", error.msg);
		label.css("color", "red")
		label.parent().tooltip('show')
		setTimeout(function () {
			label.parent().tooltip("destroy")
		}, 5000);
	}
}

function Post(url, data, callback) {
	'use strict';
	$.ajax({
		type: 'post',
		url: url,
		data: JSON.stringify(data),
		processData: false,
    	contentType: 'application/json',
    	success: callback
	})
}

function Dialog(msg, cb) {
	'use strict';
	if (msg.yes) {
		vex.dialog.buttons.YES.text = msg.yes
	}
	if (msg.no) {
		vex.dialog.buttons.NO.text = msg.no
	}
	vex.dialog.confirm({
		message: msg.message,
	  	callback: cb
	});
}

function Screen(url, afterOpen) {
	'use strict';
	$.get(url, function (res) {
		vex.open({
			contentCSS: { 'margin-top' : '-150px'},
			content: res,
			afterOpen: afterOpen
		});
	});
}

function CloseScreen() {
	vex.close()
}