//JQUERY request handlers
$(document).ready(function () {
  if (!localStorage.getItem("token")) {
    window.location.href = "/index.html";
  }
});

// file selector change label
$("#file-upload").change(function () {
  var file = $("#file-upload")[0].files[0].name;
  $("#file-upload-label").text(file);
});

// common response handler
function common_response_handler(data) {
  if (data.is_success) {
    localStorage.setItem("token", data.token);
  } else {
    alert(data.msg);
    if (new RegExp("\\blogin\\b", "gi").test(data.msg)) {
      localStorage.removeItem("token");
      window.location.href = "/index.html";
    }
  }
}

// transcribe result handler
function transcribe_result_handler(data) {
  common_response_handler(data);
  if (data.is_success) {
    alert("Transcribed text: " + data.data);
  }
}

// search result handler
function search_result_handler(data) {
  common_response_handler(data);
  if (data.is_success) {
    $("#table_data_body").html("");

    // fill table view with data
    data.data.forEach(function (item) {
      $("#table_data_body").append(
        "<tr><td>" +
          item.original_audio_file_name +
          "</td><td>" +
          item.original_audio_file_path +
          "</td><td>" +
          item.text +
          "</td></tr>"
      );
    });
  }
}

// prepare file upload request
function upload_handler(id) {
  data = new FormData($(id)[0]);
  data.set("is_save_file", $("#file-save-check").is(":checked"));
  return data;
}

//JQUERY request handlers
$(document).ready(function () {
  prepareEndpointCall_results();
});

function prepareEndpointCall_results() {
  // add file-upload and transcribe endpoint
  subscribeAjaxEvent(
    "#file-submit",
    "/transcribe",
    "form#file_upload_form",
    transcribe_result_handler,
    upload_handler,
    true
  );

  // add filter endpoint
  subscribeAjaxEvent(
    "#filter_btn",
    "/filter",
    "form#filter_form",
    search_result_handler,
    function () {
      return {
        page_no: 1,
        query: $("#search_query").val()
      };
    }
  );

  // add all-data endpoint
  subscribeAjaxEvent(
    "#get_all_data",
    "/all-data",
    "form#register-form",
    search_result_handler,
    function () {
      return {
        page_no: 1,
      };
    }
  );
}
