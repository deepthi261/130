document.getElementById('btn').onclick = function() {
	// AJAX call to server
	var xhttp = new XMLHttpRequest();
	xhttp.onreadystatechange = function() {
		// Call-back
		if (xhttp.readyState == 4 && xhttp.status == 200) {
			var message;
			if (xhttp.responseText.includes('true')) {
				message = 'User already exists';
			} else {
				message = 'New user is registered';
			}
			// Setting the error message
			document.getElementById("errorMessage").innerHTML = message;
		}
	};
	xhttp.open("POST", "isUser", true);
	var data = new FormData();
	// Setting the data on the request to server
	data.append('new-word', document.getElementById("entry").value);
	xhttp.send(data);
}