function returnUserChamp(){
		$.ajax({
	       type: 'GET',
	       dataType: 'text',
	       url: 'http://localhost:8080/players/search/all',
	        success: function (data){
	        	var array = JSON.parse(data);
	        	var tamanho = array.length;
	        	var x = document.getElementById("playerChamp");
	        	var y = document.getElementById("adminChamp");

	        	$("#playerChamp").html('<option value="" disabled selected></option>');
	        	$("#adminChamp").html('<option value="" disabled selected></option>');

	        	for(var i = 0; i < tamanho; i++){
	        		var option = document.createElement("option");
	        		option.text = array[i].firstname;
	        		option.value = array[i].cpf;
	        		x.add(option);
	        	}

	        	for(var i = 0; i < tamanho; i++){
	        		var option2 = document.createElement("option");
	        		option2.text = array[i].firstname;
	        		option2.value = array[i].cpf;
	        		y.add(option2);
	        	}
        	
        	}

    	});
}

function registerChamp(){
	var name = $("#formChamp #name").val(); 
	var players = $("#formChamp #playerChamp").val(); 
	var idate = $("#formChamp #dateChamp").val(); 
	var sets = $("#formChamp #set").val(); 
	var admin = $("#formChamp #adminChamp").val();

	var numberSets = parseInt(sets);

	var playerName = document.getElementById("playerChamp").options;

	var adminName = document.getElementById("adminChamp").options;

	var arrayPlayer = [];

	for(var j = 0; j < adminName.length; j++){
		if(adminName[j].selected == true){
			var adminTextName = adminName[j].text;
		}
	}
 

	for(var i = 0; i < playerName.length; i++){
		if(playerName[i].selected == true){
			arrayPlayer.push({"firstname":playerName[i].text, "cpf":playerName[i].value});
		}
	}

	var jsonChamp = ({"name": name, "players": arrayPlayer , "idate": idate, "sets": numberSets, "admin": {"firstname": adminTextName, "cpf": admin}});
	    $.ajax({
	       type: 'POST',
	       data: JSON.stringify(jsonChamp),
	       url: 'http://localhost:8080/championships/register',
	       dataType: 'text',
	       success: function(data) {
	        	console.log(data);
      		}
		});
}

function returnChamp(){
		$.ajax({
	       type: 'GET',
	       dataType: 'text',
	       url: 'http://localhost:8080/championships/search/all',
	        success: function (data){
	        	var array = JSON.parse(data);
	        	console.log(array);
	        	var tamanho = array.length;
	        	var table = document.getElementById("tableChamp");
	        	var arrayNome = [];
	        	$("#bodyTableChamp").html('');

	        	for(var i = 0; i < tamanho; i++){
	        		var row = table.insertRow(i);
	        		var cell1 = row.insertCell(0);
	        		var cell2 = row.insertCell(1);
	        		var cell3 = row.insertCell(2);
	        		var cell4 = row.insertCell(3);

	        		var teste = array[i].players.length;

	        		cell1.innerHTML = array[i].name;
	        		cell2.innerHTML = array[i].initialdate;
	        		cell3.innerHTML = array[i].admin.firstname;
	        		for(var j = 0; j < teste; j++){
	        			arrayNome.push(array[i].players[j].firstname);
	        		}
	        		cell4.innerHTML = arrayNome;
	        	}
        	
        	}

    	});
}