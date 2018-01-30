$("#mo").click(function(){
    function  aa(){
	var obj = {};
	obj.option = "";
	obj.input = $('#sea1').val();
	obj.type = ""
	obj.f = $("#sel2").val();
	obj.w = $("#w").val().toString();
	$('#checkMe input:checked').each(function() {
	    obj.option = obj.option+$(this).attr('name');
	});
	$('#sel1 input:checked').each(function() {
	    if (obj.type == ""){
		obj.type = $(this).attr('name');
	    }else
	    {
		obj.type = obj.type+":"+$(this).attr('name');
	    }
	});
	return JSON.stringify(obj);
    }
    $.ajax({
	url : 'http://gopiko.fr/toilet/oauth', // La ressource ciblée
	type : 'POST' ,// Le type de la requête HTTP.
	data :aa(),
	dataType : 'json',
	success : function(code_html, statut){
	    $("#jolie").html(code_html)
	},
	error : function(resultat, statut, erreur){
	    console.log(erreur)
	},

	complete : function(resultat, statut){

	}
    });
});
