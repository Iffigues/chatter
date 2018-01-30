$("#mo").click(function(){
    function  aa(){
	var obj = {};
	obj.gender = $('#sel2').val();
	obj.nbr = $('#sel1').val();
	return JSON.stringify(obj);
    }
    $.ajax({
	url : 'http://gopiko.fr/rig/oauth', // La ressource ciblée
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
