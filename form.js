const form = document.forms[0];
form.addEventListener("submit", submitFunc);
function submitFunc(e) {
    e.preventDefault();
    const url = form.action;
    const loginDetails = {};
    loginDetails.email = form.useremail.value;
    loginDetails.password = form.userpassword.value;
    console.log(loginDetails);
    fetch(url, {
        body: JSON.stringify(loginDetails),
        method: "POST",
        mode: "no-cors",
        headers: {
            "Content-Type": "application/json",
        }
        })
        .then(res => res.json())
        .then(data => {
            console.log(data);
            // alert("success")
        }).catch(error => {
            console.log(error);
            alert("error")
        })
    }