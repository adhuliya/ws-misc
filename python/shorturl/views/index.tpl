<!DOCTYPE html> <html lang="en"> <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="description" content="utility to shorten the url">
    <meta name="keywords" content="shorten url, url shortner">
    <meta name="author" content="Anshuman Dhuliya">

    <!-- set shortcut icon when needed -->
    <link rel="shortcut icon" href="/static/image/icons8-network-care-64.png" type="image/png">

    <!-- font-awesome.min.css remote and local link -->
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
    <!-- <link rel="stylesheet" href="/static/css/font-awesome.min.css" /> -->

    <!-- w3.css remote and local link -->
    <!-- <link rel="stylesheet" href="https://www.w3schools.com/w3css/4/w3.css"> -->
    <link rel="stylesheet" href="/static/css/w3.css" />
    <link rel="stylesheet" href="/static/css/main.css" />

    <title>Shorten URL</title>

  </head><body class="w3-content">

    <h1 class="w3-padding-16 w3-text-teal w3-center">{{user.name}}</h1>

    <div class="w3-margin">
      <input class="w3-input" id="url" name="url" placeholder="Paste URL here e.g. https://www.google.com/" type="text" />
    </div>
    <div class="w3-margin w3-center">
      <button class="w3-button w3-red" id="shorten" type="button">Shorten URL</button>
    </div>

    % if len(user.urls) == 0:
    <div class="w3-container w3-center w3-margin">
    <p>You have no URLs :(</p>
    </div>
    % end

    % if len(user.urls) > 0:
    <div class="w3-margin">
    <table class="w3-table-all w3-striped" style="width:100%">
    % for url in reversed(user.urls):
      <tr id="tr-{{url}}"><td>
          <div class="w3-bar w3-container">
            <div class="w3-bar-item w3-left w3-text-teal">
              <b><a href="/{{url}}">{{url}}</a></b>
            </div>

            <div class="w3-bar-item w3-left">
              <a id="a-{{url}}" href="{{url_map[url]}}">{{url_map[url][:40]}}</a>
            </div>

            <!-- Trigger/Open the Modal -->
            <div class="w3-bar-item w3-right">
              <span onclick="delete_url('{{url}}');"
                    class="w3-button w3-round-xlarge"
                    title="Delete"><b class="fa fa-trash" style="font-size:1.5em;"></b></span>
              <span onclick="document.getElementById('modal02').style.display='block'"
                    class="w3-button w3-round-xlarge"
                    title="Edit"><b class="fa fa-pencil" style="font-size:1.5em;"></b></span>
            </div>
          </div> 
        </td></tr>
    % end
    </table>
    </div>
    % end

    <footer class="w3-container w3-padding-16 w3-text-gray w3-small">
      <br /> <br /> <br />
      <div class="w3-center">
        <p>shorten url service</p>
        <h6 class="w3-padding-16 w3-text-teal w3-center">its<span class="w3-text-deep-orange">of</span>life</h6>
      </div>
    </footer>

    <!-- extra page elements -->
    <!-- The Modal -->
    <div id="modal01" class="w3-modal">
      <div class="w3-modal-content" style="max-width:40em;">
        <div class="w3-container">
          <span onclick="document.getElementById('modal01').style.display='none'"
                class="w3-button w3-display-topright">&times;</span>
          <div class="w3-container w3-padding">
          <h4><i class="fa fa-trash"></i> Delete?</h4>
          <p id="del_short_url" class="w3-text-teal w3-large">Short URL</p>
          <p id="del_full_url">Full URL</p>
          <button id="del_url_btn" class="w3-button w3-red" type="button">Delete</button>
          </div>
        </div>
      </div>
    </div>

    <div id="modal02" class="w3-modal">
      <div class="w3-modal-content" style="max-width:40em;">
        <div class="w3-container">
          <span onclick="document.getElementById('modal02').style.display='none'"
                class="w3-button w3-display-topright">&times;</span>
          <div class="w3-container w3-padding">
        <form id="form_register" action="/u/login_check" method="post">
          <input class="w3-input w3-margin" type="text" name="edit_full_url" placeholder="Full URL" required="" />
          <input class="w3-input w3-margin" type="text" name="edit_short_url" required="" />

          <button class="w3-button w3-teal" type="submit">Modify</button>
          <button class="w3-button w3-teal" type="reset">Reset</button>
        </form>
        </div>
        </div>
      </div>
    </div>

    <!-- jquery remote and local link -->
    <!-- <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script> -->
    <script src="/static/js/jquery.min.js"></script>
    <script>
    function delete_url(short_url) {
      console.log("delete_url");
      $("#del_short_url").html(short_url);
      $("#del_full_url").html($("#a-"+short_url).attr("href"));
      $("#del_url").attr("href", "/u/del/"+short_url);
      $('#modal01').show(); // .hide();

      $("#del_url_btn").on('click', function () {
        $.ajax({
          url: '/u/del/'+short_url,
          type: 'get',
          dataType: 'json',
          contentType: 'application/json',
          success: function (data) {
            console.log('deleted');
            $('#tr-'+short_url).hide();
          },
          error: function (data) {
            console.log('error');
          },
          data: ""
        });
        $("#modal01").hide();
      });
    }

    $("#shorten").on('click', function () {
     var value = $("#url").val()
     alert("."+value+".");
     value = $.trim(value);
     if (value.length) {
       alert("."+value+".");
     }
    });
    </script>

  </body> </html>
