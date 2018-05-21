$(".fafa").on("click",function(){
    var t = this.id;
    var l = t.split("-");
    var ii = "#ml-"+l[1]
    $(".chata").hide();
    $(ii).show();
})

$("#chatcreate").on("click",function(){
    function  aa(){
	alert("eees")
	var obj = {};
	obj.name = $("#sea1").val();
	if ($("#sea2").val()!=""){
	    obj.password = $("#sea2").val();
	}
	return JSON.stringify(obj);
    }
    $.ajax({
	url : 'http://gopiko.fr/room/create', // La ressource ciblée
	type : 'POST' ,// Le type de la requête HTTP.
	data :aa(),
	success : function(code_html, statut){
	    console.log(statut)
	},

	error : function(resultat, statut, erreur){
	    alert("sss")
	    alert(erreur)
	},
	complete : function(resultat, statut){
	}
    });
    
})


$("#sea3").on("click",function(){
    function  aa(){
	var obj = [];
	obj.Search = ("#chatrr").value
	return JSON.stringify(obj);
    }

    function jsonroom(a){
	console.log(a)
	for (i = 0; i < a.length; i++){
	    var div = document.createElement('div');
	    div.className = "card"
	    var body = document.createElement('div');
	    body.className = "card-body"
	    var p = document.createElement('p');
	    var d = document.createElement('p');
	    var f = document.createElement('p');
	    
	    p.innerHTML = a[i].name
	    body.appendChild(p)
	    d.innerHTML = a[i].private
	    body.appendChild(d)
	    f.innerHTML = a[i].number
	    body.appendChild(f)
	    
	    div.appendChild(body)
	    
	    document.getElementById('dock').appendChild(div);
	}
    }
    
    $.ajax({
	url : 'http://gopiko.fr/room/search', // La ressource ciblée
	type : 'POST' ,// Le type de la requête HTTP.
	data :aa(),
	success : function(code_html, statut){
	    jsonroom(code_html)
	},

	error : function(resultat, statut, erreur){
	    alert(erreur)
	},
	complete : function(resultat, statut){
	}
    });

})


$("#mo").click(function(){
    function  aa(){
	var obj = {};
	obj.Arg = "";
	$('#checkMe input:checked').each(function() {
	   obj.Arg = obj.Arg+ $(this).attr('name');
	});
	obj.Input = $('#sea1').val();
	obj.M = $("#m").val();
	obj.Input2 = $("#sea2").val();
	return JSON.stringify(obj);
    }
    $.ajax({
	url : 'http://gopiko.fr/fortune/oauth', // La ressource ciblée
	type : 'POST' ,// Le type de la requête HTTP.
	data :aa(),  
	dataType : 'json',
	success : function(code_html, statut){
	  
	    $("#fr").text(code_html)
	},

	error : function(resultat, statut, erreur){
	    alert(erreur)
	},

	complete : function(resultat, statut){

	}
    });
});
