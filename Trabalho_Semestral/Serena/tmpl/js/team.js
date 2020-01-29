function returnUserTeam(){
		$.ajax({
	       type: 'GET',
	       dataType: 'text',
	       url: 'http://localhost:8080/players/search/all',
	        success: function (data){
	        	var array = JSON.parse(data);
	        	var tamanho = array.length;
	        	var x = document.getElementById("playersTeamOptionOne");
	        	var y = document.getElementById("playersTeamOptionTwo");

	        	$("#playersTeamOptionOne").html('<option value="" disabled selected></option>');
	        	$("#playersTeamOptionTwo").html('<option value="" disabled selected></option>');

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

function registerTeam() {
	var nameTeam = $("#formTeam #name").val(); 
	var playerOne = document.getElementById("playersTeamOptionOne").options;
	var playerTwo = document.getElementById("playersTeamOptionTwo").options;

	for(var i = 0; i < playerOne.length; i++){
		if(playerOne[i].selected == true){
			var name = playerOne[i].text;
			var value = playerOne[i].value;
		}
	}

	for(var i = 0; i < playerTwo.length; i++){
		if(playerTwo[i].selected == true){
			var name2 = playerTwo[i].text;
			var value2 = playerTwo[i].value;
		}
	}
	
	var jsonTeam = ({"name": nameTeam, "playerone": {"firstname": name, "cpf": value}, "playertwo": {"firstname": name2, "cpf": value2}});

	    $.ajax({
	       type: 'POST',
	       data: JSON.stringify(jsonTeam),
	       url: 'http://localhost:8080/team/register',
	       dataType: 'text',
	       success: function(data) {
	        	console.log(data);
    		}
		});
}

function returnTeam(){
		$.ajax({
	       type: 'GET',
	       dataType: 'text',
	       url: 'http://localhost:8080/team/search/all',
	        success: function (data){
	        	var array = JSON.parse(data);
	        	var table = document.getElementById("tableTeam");

	        	console.log(array);
	
	        	for(var i = 0; i < array.length; i++){
	        		var row = table.insertRow(i);
	        		var cell2 = row.insertCell(0);
	        		var cell1 = row.insertCell(1);

	        		cell2.innerHTML = array[i].player1.firstname+" e "+array[i].player2.firstname;
	        		cell1.innerHTML = array[i].name;
	        	}
        	
        	}

    	});
}