$("#mo").click(function(){
    function  aa(){
	var obj = {};
	obj.help = [];
	obj.style = "";
	obj.input = $('#sea1').val();
	obj.type = $("#sel1").val();
	obj.animal = $("#sel2").val();
	obj.eyes = $("#e").val();
	obj.tong = $("#T").val();
	obj.col = $("#W").val().toString();
	$('#checkMe input:checked').each(function() {
	    obj.help.push($(this).attr('name'));
	});
	$('#checkMoi input:checked').each(function() {
	    obj.style = obj.style+$(this).attr('name');
	});
	return JSON.stringify(obj);
    }
    $.ajax({
	url : 'http://gopiko.fr/cowsay/oauth', // La ressource ciblée
	type : 'POST' ,// Le type de la requête HTTP.
	data :aa(),
	dataType : 'json',
	success : function(code_html, statut){
	    $("#comment").text(code_html)
	},
	error : function(resultat, statut, erreur){
	    console.log(erreur)
	},

	complete : function(resultat, statut){

	}
    });
});
