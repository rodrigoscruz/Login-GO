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

	fetch("./register",{

		headers: new Headers({ "X-Content-Type-Options": "multipart/form-data",}),
		method: "POST",
		body: new FormData(form),
	})
	.then((response) => {

		const responseStatus = {

			200:() => { alert("Dados enviados com sucesso.");},
			400:() => { alert("Este cadastro já existe.");},
			404:() => { alert("Tente realizar o cadastro mais tarde.");}		
		}
	

		if (responseStatus[response.status]) {

			let responseUser = responseStatus[response.status];
			responseUser();
		}else{
			alert("Realize o cadastro nocamente.");
		}
	})

}

