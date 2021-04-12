const signUpButton = document.getElementById("signUp");
const signInButton = document.getElementById("signIn");
const container = document.getElementById("container");
const endpoint_root = "http://localhost:5000";

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
    alert("success!");
  } else {
    alert(data.msg);
  }
}

//JQUERY request handlers
$(document).ready(function () {
  prepareEndpointCall();
});

function prepareEndpointCall() {
  callAjax("#signIn_btn", "/login", "form#login-form", logged);
  callAjax("#register_btn", "/register", "form#register-form", logged);
}

function callAjax(elm_id, endpoint, from_id, callback) {
  $(elm_id).click(function () {
    $.ajax({
      url: endpoint_root + endpoint,
      type: "post",
      dataType: "json",
      data: $(from_id).serialize(),
      complete: function (data) {
        callback(data.responseJSON);
      },
    });
  });
}
