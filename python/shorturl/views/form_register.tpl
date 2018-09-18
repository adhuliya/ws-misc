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

    <a href="/" style="text-decoration:none;"><h1 class="w3-padding-16 w3-text-teal w3-center">its<span class="w3-text-deep-orange">of</span>life</h1></a>

    <div class="w3-content w3-container w3-card-4" style="max-width:50em;">
      <div class="w3-container w3-panel">
        <h2>Register</h2>
        <form id="form_register" action="/u/login_check" method="post">
          <input class="w3-input w3-margin" type="text" name="name" placeholder="Your Name" required="" />
          <input class="w3-input w3-margin" type="text" name="email_id" placeholder="Email ID (will work as User Id)" required="" />
          <input class="w3-input w3-margin" type="password" name="passwd1" placeholder="Password" required="" />
          <input class="w3-input w3-margin" type="password" name="passwd2" placeholder="Retype Password" required="" />

          <button class="w3-button w3-teal" type="submit">Register</button>
          <button class="w3-button w3-teal" type="reset">Reset</button>
        </form>
      </div>
    </div>

    <footer class="w3-container w3-padding-16 w3-text-gray w3-small">
      <br /> <br /> <hr class="w3-content" style="max-width:60%;"/>
      <div class="w3-center">
        <p>shorten url service</p>
        <h6 class="w3-padding-16 w3-text-teal w3-center">its<span class="w3-text-deep-orange">of</span>life</h6>
      </div>
    </footer>

    <!-- jquery remote and local link -->
    <!-- <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script> -->
    <script src="/static/js/jquery.min.js"></script>

  </body> </html>
