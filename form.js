const form = document.forms[0];
form.addEventListener("submit", submitFunc);
function submitFunc(e) {
    e.preventDefault();
    // const url = form.action;
    const base_url = "http://localhost:10000/login"
    const loginDetails = {};
    loginDetails.email = form.useremail.value;
    loginDetails.password = form.userpassword.value;
    console.log(loginDetails);
    fetch(base_url, {
        body: JSON.stringify(loginDetails),
        method: "POST",
        //credentials: 'same-origin',
        mode: "no-cors",
        headers: {
            // Accept: 'application/json',
            "Content-Type": "application/json",
        },
        })
        .then(res => {if (!res.ok) {
            throw `Server error: [${res.status}] [${res.statusText}] [${res.url}]`;
        }
        return res.json()})
        .then(data => {
            console.log(data);
            // alert("success")
        })
        .catch(error => {
            console.log(error);
            alert(error)
        })
        // const getForm = async() => {
        //     const response = await fetch("http://127.0.0.1:8080/login", {
        //         method: 'POST',
        //         body: JSON.stringify(loginDetails),
        //         mode:"no-cors",
        //         headers: {
        //             "Content-Type": "application/json",
        //         }
        //       });
        //       const string = await response.text();
        //       const json = string === "" ? {} : JSON.parse(string);
        //     //   return json;
        //       console.log(json);
        // }
        // getForm();
    }