const endpoint_root = "";
// const endpoint_root = !new RegExp(":5000").test(window.location.host)
//   ? window.location.protocol + "//" + window.location.hostname + ":5000"
//   : "";

// ajax request handler
function subscribeAjaxEvent(
  elm_id,
  endpoint,
  from_id,
  callback_fn,
  custom_value_fn = null,
  is_file_upload = false
) {
  $(elm_id).click(function () {
    // input validation
    if (
      $(from_id)
        .find("input[name]")
        .filter(function () {
          return $(this).val().length < 1;
        }).length > 0
    ) {
      alert("Please fill all the required fields!");
      return false;
    } else {
      // validation is ok. now do the ajax call.
      let addToken = (data, f) => {
        if (f) {
          data.set("token", localStorage.getItem("token"));
        } else {
          data.token = localStorage.getItem("token");
        }
        return data;
      };

      // disable the submit button
      $(elm_id).prop("disabled", true);

      let options = {
        url: endpoint_root + endpoint,
        type: "POST",
        enctype: "multipart/form-data",
        dataType: "json",
        data: custom_value_fn
          ? addToken(custom_value_fn(from_id), is_file_upload)
          : $(from_id).serialize(),

        cache: false,
        complete: function (data) {
          // enable the submit button
          $(elm_id).prop("disabled", false);
          callback_fn(data.responseJSON);
        },
        //xhrFields: { withCredentials: true },
      };

      // set processing off fro file upload
      if (is_file_upload) {
        options.processData = false;
        options.contentType = false;
      }

      $.ajax(options);
    }
  });
}
