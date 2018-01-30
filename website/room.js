function addRooming(name,nbr,privates){
    var type = "";
    if (privates == true )
    {
	type = "private";
    }else{
	type = "public";
    }
    var cont = '<div class="container-fluid>"';
    var row = '<div class="row">';
    var col = '<div class="col-sm-6">';
    var end = '</div>';
    var input = '<input type="text" placeholder="password"><input>'

    var naming = col+"Name"+end+col+name+end;
    var number = col+"number"+end+col+nbr+end
    var types = col+"type"+end+col+type+end
    var password = col+"password"+end+col+nbr+end
}
