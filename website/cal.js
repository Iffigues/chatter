$("#mo").click(function(){
    function  aa(){
	var obj = {};
	obj.A = $('#A').val();
	obj.B = $('#B').val();
	obj.W = $('#W').val();
	obj.T = $('#t').val();
	return JSON.stringify(obj);
    }
    $.ajax({
	url : 'http://gopiko.fr/calendar/oauth', // La ressource ciblée
	type : 'POST' ,// Le type de la requête HTTP.
	data :aa(),
	dataType : 'json',
	success : function(code_html, statut){
	    console.log(code_html)
	    $("#jolien").text(code_html)
	},
	error : function(resultat, statut, erreur){
	    console.log("eee")
	    console.log(erreur)
	},

	complete : function(resultat, statut){

	}
    });
});
