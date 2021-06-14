const form = document.forms[0];
form.addEventListener("submit", submitFunc);

function submitFunc(e) {
    const url = "http://localhost:8080/login";
    const loginDetails = {};

    loginDetails.email = form.useremail.value;
    loginDetails.password = form.userpassword.value;

    const fetchOptions = {
        body: JSON.stringify(loginDetails),
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        }
    };

    fetch(url, fetchOptions)
    .then(response => response.json())
    .then(data =>  {
        alert(data.message);
    })
    .catch(error => console.log(error));

    e.preventDefault();

}