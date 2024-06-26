function validate() {
  let pass = document.getElementById("pass").value;
  let conf = document.getElementById("conf").value;
  let error = document.getElementById("error");
  error.innerHTML = "";

  let but = document.getElementById("subm");
  console.log(pass, conf);
  if (pass !== conf) {
    error.innerHTML = "Passwords do not match";
    but.classList.replace("Ben", "Bis");
    return;
  }
  if (pass.match("^\\s+$")) {
    error.innerHTML = "Password cannot be empty";
    but.classList.replace("Ben", "Bis");
    return;
  }
  if (pass === "") {
    but.classList.replace("Ben", "Bis");
    return;
  }
  if (pass.legth < 8) {
    error.innerHTML = "Password has to be longer than 8 char";
    but.classList.replace("Ben", "Bis");
    return;
  }
  but.classList.replace("Bis", "Ben");
}
