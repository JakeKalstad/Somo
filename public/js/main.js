var User = {};

function LoadUser (uid, callback) {
	'use strict';
	Post('/user/get', {id: uid}, function (res) {
		User = res;
		callback(User);
	});
}

function artistDeclined() {
	'use strict';
}

function SaveUser(cb) {
	'use strict';
	cb = cb || function(){};
	var trax = User.artist.tracks;
	Post('/user/save', User, function (res) {
		cb(res)
	});
	User.artist.tracks = trax;
}

function createArtist() {
	'use strict';
	Screen('/user/createartist',  function($vexContent) {
		function handleFileSelect(evt) {
			var files = evt.target.files;
			var f = files[0];
			var reader = new FileReader();
			reader.onload = (function (theFile) {
				return function (e) {
					var span = document.createElement('span');
					span.innerHTML = ['<img class="artist_img" src="', e.target.result, '" title="', escape(theFile.name), '"/>'].join('');
					document.getElementById('img_placeholder').insertBefore(span, null);
				};
			})(f);
			reader.readAsDataURL(f);
		}
		document.getElementById('my_img_upload_input').addEventListener('change', handleFileSelect, false);
		$('.artist_create_btn').on('click', function () {
			User.artist = {
				moniker: $(".artist_name").val(),
				img_url: $(".artist_img").attr("src"),
				blurb: $(".artist_blurb").val(),
			}
			SaveUser();
		});
	})
}
function ReloadStream() {
	Post('/audio/stream', User, function (res) {
			$('.somo_content').html(res)
			$('.audio .delete').click(function () {
				for (var i=0; i < User.artist.tracks.length; i++) {
				    if (User.artist.tracks[i].id == $(this).data('id')) {
				        User.artist.tracks.splice(i, 1);
				    }
				}
				SaveUser(function () {
					ReloadStream();
				});				
			});
	});
}
var LoadFeed = (function () {
	'use strict';
	function handlers() {
		$('#upload_btn').click(function () {
			LoadUploadTrack();
		});
		$("#account_btn").click(function () {
			LoadAccount();
		});
	}
	function updateUi() {
		$('#upload_btn').removeClass('hidden');
		$('#account_btn').removeClass('hidden');
	}
	if (User.artist.tracks && User.artist.tracks.length > 0) {
		ReloadStream()
	} else {

	}
	handlers();
	updateUi();
});

var LoadDashboard = (function () {
	'use strict';
	if (User.artist.denied) {
		LoadFeed()
	} else if (User.artist.id == "") {
		Dialog({
			yes: "Sounds Grooovy", 
			no: "Nope, I wanna listen.", 
			message: 'Would you like an artist account to upload sweet beets?'
		}, function (value) {
			if (value) {
				createArtist();
			} else {
				artistDeclined();
			}
		});
	} else {
		LoadFeed()
	}
});

var LoadNewTour = (function () {

});

var LoadAccount = (function () {
	Screen('/user/account', function ($vexContent) {

	});
});

var LoadUploadTrack = (function () {
	'use strict';
	Screen('/audio/upload', function ($vexContent) {
		var track = {
			file: 0,
			name: 0,
			description: 0,
			tags: 0
		};
		$('#file1').change(function(evt, res) {
			var file = this.files[0];
			var name = file.name;
			var size = file.size;
			var type = file.type;
		});
		function progressHandlingFunction(e){
			if(e.lengthComputable){
		       // $('progress').attr({value:e.loaded,max:e.total});
		   }
		}
		function getTags() {
			return ['hip-hop', 'trap', 'crack-rock-steady'];
		}
		$(".track_upload.btn-success").on('click', function(e) {
			e.preventDefault();
			var self = $("#track_upload_form")[0];
			var formData = new FormData(self);
			$.ajax({
				url: self.action,
				data: formData,
				type: "POST",
				dataType: "json",
				xhr: function() {  
					var myXhr = $.ajaxSettings.xhr();
					if(myXhr.upload){ 
		                myXhr.upload.addEventListener('progress', progressHandlingFunction, false); // For handling the progress of the upload
		            }
		            return myXhr;
		        },
		        success: function (res) {
		        	var track = {
		        		name: $($("#track_upload_form")[0].track_name).val(),
		        		description: $($("#track_upload_form")[0].track_description).val(),
		        		tags: getTags(),
		        		handle: res.Files[0].handler,
		        		filename: res.Files[0].Filename
		        	}
		        	if (!User.artist.tracks) {
		        		User.artist.tracks = [];
		        	}
		        	User.artist.tracks.push(track);
		        	SaveUser()
		        	ReloadStream()
		        	CloseScreen()
		        },
		        cache: false,
		        contentType: false,
		        processData: false
		    });
			return false;
		});
});
});

var LoadSignUp = (function () {
	'use strict';
	Screen('/user/signup', function($vexContent) {
		$('.form').find('input, textarea').on('keyup blur focus', function (e) {
			var $this = $(this),
			label = $this.prev('label');
			if (e.type === 'keyup') {
				if ($this.val() === '') {
					label.removeClass('active highlight');
				} else {
					label.addClass('active highlight');
				}
			} else if (e.type === 'blur') {
				if( $this.val() === '' ) {
					label.removeClass('active highlight'); 
				} else {
					label.removeClass('highlight');   
				}   
			} else if (e.type === 'focus') {
				if( $this.val() === '' ) {
					label.removeClass('highlight'); 
				} 
				else if( $this.val() !== '' ) {
					label.addClass('highlight');
				}
			}
		});
		$('.tab a').on('click', function (e) {
			e.preventDefault();
			$(this).parent().addClass('active');
			$(this).parent().siblings().removeClass('active');
			var target = $(this).attr('href');
			$('.tab-content > div').not(target).hide();
			$(target).fadeIn(600);
		});
		$('.create_button').click(function (e) {
			e.preventDefault()
			Post('/User/Save', {
				login: $('.s_login').val(),
				password: $('.s_password').val(),
				email: $('.s_email').val()						
			}, function (res) {
				if (!res.errors) {
					User = res.user;
					LoadNewTour();
				} else {
					displayErrors(res.errors);
				}
			});
			return true
		});
		$('.sign_in_button').click(function (e) {
			e.preventDefault()
			Post('/User/Auth', {
				login: $('.l_login').val(),
				password: $('.l_password').val()						
			}, function (res) {
				if (res.authorized) {
					vex.close()
					User = res.user;
					LoadDashboard();
				} else {
					displayErrors(res.errors);
				}
			});
			return true
		});
	});
});

$(document).ready(function () {
	async.parallel([function (cb) {
		LoadUser($('#user_info').data('user-id'), function (res) {
			cb();
		});
	}, function (cb) {
		cb();
	}], function () {
		if(User.id == "") {
			LoadSignUp();
		} else { 
			LoadDashboard();
		}
	});
});