$("#mo").click(function(){
    function  aa(){
	var obj = {};
	obj.option = "";
	obj.input = $('#sea1').val();
	obj.f = $('#sel1').val();
	obj.w = $("#w").val().toString();
	$('#checkMe input:checked').each(function() {
	    obj.option = obj.option+$(this).attr('name');
	});
	return JSON.stringify(obj);
    }
    $.ajax({
	url : 'http://gopiko.fr/figlet/oauth', // La ressource ciblée
	type : 'POST' ,// Le type de la requête HTTP.
	data :aa(),
	dataType : 'json',
	success : function(code_html, statut){
	    $("#jolie").text(code_html)
	},
	error : function(resultat, statut, erreur){
	    console.log(erreur)
	},

	complete : function(resultat, statut){

	}
    });
});
