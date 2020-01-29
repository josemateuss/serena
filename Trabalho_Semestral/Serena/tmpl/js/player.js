function registerUser() {

	var firstname = $("#formPlayer #firstname").val(); 
	var cpf = $("#formPlayer #cpf").val(); 
	var phone = $("#formPlayer #phone").val(); 
	var email = $("#formPlayer #email").val(); 
	var password = $("#formPlayer #pass").val(); 
	var birthdate = $("#formPlayer #birthdate").val(); 
	var class1 = $("#formPlayer #class").val(); 
	var group = $("#formPlayer #group").val(); 
	var category = $("#formPlayer #category").val(); 

	
	var jsonPlayer = ({"firstname": firstname, "cpf": cpf, "phone": phone, "email": email, "pass": password, "birthdate": birthdate, "class": class1, "group": group, "category": category});

	    $.ajax({
	       type: 'POST',
	       data: JSON.stringify(jsonPlayer),
	       url: 'http://localhost:8080/players/register',
	       dataType: 'text',
	       success: function(data) {
	        	console.log(data);
    		}
		});
}

function returnUser(){
		$.ajax({
	       type: 'GET',
	       dataType: 'text',
	       url: 'http://localhost:8080/players/search/all',
	        success: function (data){
	        	var array = JSON.parse(data);
	        	var table = document.getElementById("tablePlayer");
	        	$("#bodyTablePlayer").html('');
	
	        	for(var i = 0; i < array.length; i++){
	        		var row = table.insertRow(i);
	        		var cell4 = row.insertCell(0);
	        		var cell3 = row.insertCell(1);
	        		var cell2 = row.insertCell(2);
	        		var cell1 = row.insertCell(3);

	        		cell4.innerHTML = array[i].firstname;
	        		cell3.innerHTML = array[i].group;
	        		cell2.innerHTML = array[i].class;
	        		cell1.innerHTML = array[i].category;
	        	}
        	
        	}

    	});
}