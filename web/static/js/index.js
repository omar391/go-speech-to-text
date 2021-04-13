const signUpButton = document.getElementById("signUp");
const signInButton = document.getElementById("signIn");
const container = document.getElementById("container");


signUpButton.addEventListener("click", () => {
  container.classList.add("right-panel-active");
});

signInButton.addEventListener("click", () => {
  container.classList.remove("right-panel-active");
});

// endpoint result handlers
function logged(data) {
  if (data.is_success) {
    localStorage.setItem("token", data.token);
    window.location.href = '/results.html'
  } else {
    localStorage.removeItem("token");
    alert(data.msg);
  }
}

//JQUERY request handlers
$(document).ready(function () {
  prepareEndpointCall_index();
});

function prepareEndpointCall_index() {
  subscribeAjaxEvent("#signIn_btn", "/login", "form#login-form", logged);
  subscribeAjaxEvent("#register_btn", "/register", "form#register-form", logged);
}
