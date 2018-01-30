function toImg(a,can){
    var aa = document.getElementById(a)
    var canvas = document.getElementById('canvas');
    var w = $(aa).width();
    var h = $(aa).height();
    var ctx = canvas.getContext('2d');
    //canvas.height = h;
    //canvas.width = w;
    console.log(w,h)
    var data = '<svg xmlns="http://www.w3.org/2000/svg" width="200" height="200">' +
	'<foreignObject width="100%" height="100%">' +
	'<div xmlns="http://www.w3.org/1999/xhtml" style="font-size:40px">' +
	'<em>I</em> like ' +
	'<span style="color:white; text-shadow:0 0 2px blue;">' +
	'cheese</span>' +
	'</div>' +
	'</foreignObject>' +
	'</svg>';

    var DOMURL = window.URL || window.webkitURL || window;
    var img = new Image();
    var svg = new Blob([data], {type: 'image/svg+xml'});
    var url = DOMURL.createObjectURL(svg);

    img.оnload = function() {
	ctx.drawImage(img, 0, 0);
	DOMURL.revokeObjectURL(url);
    }

    img.src = url
    
}
