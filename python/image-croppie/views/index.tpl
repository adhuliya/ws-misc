<!DOCTYPE html>
<!--Anshuman Dhuliya-->
<!--croppie documentation-->
<!--http://foliotek.github.io/Croppie/#documentation-->
<html lang ="en">
  <head>
    <meta charset="UTF-8" >
    <title>A working croppie example</title>
    <link rel="stylesheet" href="/static/js/Croppie-2.6.2/croppie.css" />

<style>
#page {
  background: #FFF;
  padding: 20px;
  margin: 20px;
}

#demo-basic {
  width: 200px;
  height: 300px;
}
</style>

  </head>

  <body>

<div id="page">
  <div id="demo-basic"></div>
</div>

    <div id="values"></div>
    <button id="button1" type="button" data-int="12">Hello</button>
    <button id="button2" type="button" data-int="15">Rotate</button>
    <button id="result" type="button" data-int="15">Result</button>
    <input id="upload" value="Choose a file" accept="image/*" type="file">

    <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.0/jquery.min.js"></script>
    <script src="/static/js/Croppie-2.6.2/croppie.js"></script>
    <script src="/static/js/Croppie-2.6.2/demo/demo.js"></script>

    <script>
    console.log("hello");
var el = document.getElementById('demo-basic');
var vanilla = new Croppie(el, {
    viewport: { width: 200, height: 200 },
    boundary: { width: 300, height: 300 },
    showZoomer: false,
    enableOrientation: true
});
vanilla.bind({
    url: '/static/js/Croppie-2.6.2/demo/demo-1.jpg',
    orientation: 1
});

//    var gbasic;
//    $(function() {
//      var basic = $('#demo-basic').croppie({
//        showZoomer: false,
//        enableOrientation: true,
//        viewport: {
//          width: 200,
//          height: 200,
//          type: 'square' //or 'circle'
//        }
//      });
//      basic.croppie('bind', {
//        url: 'Croppie-2.6.2/demo/demo-1.jpg',
//        points: [77, 269, 280, 739]
//      });
//      gbasic = basic;
//    });
//    $('#demo-basic').on('update.croppie', function(ev, cropData) {console.log(cropData);});

    $('#button1').on('click', function(ev) { console.log($(this).data("int"), "Clicked");});
    $('#button2').on('click', function(ev) { vanilla.rotate(90);});
		$('#upload').on('change', function () { readFile(this); });

		function readFile(input) {
 			if (input.files && input.files[0]) {
	            var reader = new FileReader();
	            
	            reader.onload = function (e) {
					$('.upload-demo').addClass('ready');
	            	vanilla.bind({
	            		url: e.target.result
	            	}).then(function(){
	            		console.log('jQuery bind complete');
	            	});
	            	
	            }
	            
	            reader.readAsDataURL(input.files[0]);
	        }
	        else {
		        swal("Sorry - you're browser doesn't support the FileReader API");
		    }
		}
		document.querySelector('#result').addEventListener('click', function (ev) {
      console.log("result clicked");
			vanilla.result({
				type: 'blob'
			}).then(function (blob) {
        console.log("hi", typeof(blob), blob);
var formData = new FormData();

  formData.append('img_blob', blob);
  formData.append('img_type', "image/png");

  $.ajax('/path/to/upload', {
    method: "POST",
    data: formData,
    processData: false,
    contentType: false,
    success: function () {
      console.log('Upload success');
    },
    error: function () {
      console.log('Upload error');
    }
  });
//				popupResult({
//					src: window.URL.createObjectURL(blob)
//				});
			});
		});
    </script>
  </body>

</html>
