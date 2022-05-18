var form = document.getElementsByTagName('form')[0];

form.addEventListener("submit", getData);

function sendData(method, url, header, data, serverResponse){

	let httpRequest = new XMLHttpRequest();

	httpRequest.open(method, url);

	httpRequest.setRequestHeader("X-Content-Type-Options", header);

	httpRequest.send(data);

	httpRequest.onreadystatechange = serverResponse;
}

function getData(event) {
	
	event.preventDefault();

	data = {
		name: form.name.value,
		lastname: form.lastname.value,
		email: form.email.value,
		password: form.password.value,
		confirm: form.confirm.value,
		accept: form.accept.value,
	}

	let json = JSON.stringify(data)

	sendData(
		"POST",
		"./account",
		"application/json",
		json,
		response
		)

}

