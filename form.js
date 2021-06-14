const form = document.forms[0]
form.addEventListener("submit", submitFunc, false)
function submitFunc(e) {
    // e.preventDefault();
    const loginDetails = {}
    loginDetails.email = form.useremail.value;
    loginDetails.password = form.userpassword.value
    alert(JSON.stringify(loginDetails))
    console.log(JSON.stringify(loginDetails));
    fetch("http://127.0.0.1:8080/processpost", {
        body: JSON.stringify(loginDetails),
        method: "POST",
        mode: "no-cors",
        headers: {
            "Content-Type": "application/json",
        }
        }
        ).then(response => {
            if(response.ok){console.log("success")
        } else{
            console.log("error");
        }
        })
        .then(response => response.json())
        .then(data => {
            console.log(data);
        });
    }
