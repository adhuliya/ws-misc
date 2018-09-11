<!DOCTYPE html>
<html>
<head>
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<style>
* {
  box-sizing: border-box;
}

body {
  font: 16px Arial;  
}

.autocomplete {
  /*the container must be positioned relative:*/
  position: relative;
  display: inline-block;
}

input {
  border: 1px solid transparent;
  background-color: #f1f1f1;
  padding: 10px;
  font-size: 16px;
}

input[type=text] {
  background-color: #f1f1f1;
  width: 100%;
}

input[type=submit] {
  background-color: DodgerBlue;
  color: #fff;
  cursor: pointer;
}

.autocomplete-items {
  position: absolute;
  border: 1px solid #d4d4d4;
  border-bottom: none;
  border-top: none;
  z-index: 99;
  /*position the autocomplete items to be the same width as the container:*/
  top: 100%;
  left: 0;
  right: 0;
}

.autocomplete-items div {
  padding: 10px;
  cursor: pointer;
  background-color: #fff; 
  border-bottom: 1px solid #d4d4d4; 
}

.autocomplete-items div:hover {
  /*when hovering an item:*/
  background-color: #e9e9e9; 
}

.autocomplete-active {
  /*when navigating through the items using the arrow keys:*/
  background-color: DodgerBlue !important; 
  color: #ffffff; 
}
</style>
</head>     

<body>

<h2>Autocomplete</h2>

<p>Start typing:</p>

<!--Make sure the form has the autocomplete function switched off:-->
<form autocomplete="off" action="/action_page.php">
  <div class="autocomplete" style="width:300px;">
    <input id="myInput" type="text" name="myCountry" placeholder="Country">
  </div>
  <input type="submit">
</form>

<script src="/static/js/family-async.js"></script>

<script>
function autocomplete(inp) {
  /*the autocomplete function takes two arguments,
  the text field element and an array of possible autocompleted values:*/
  var currentFocus;
  /*execute a function when someone writes in the text field:*/
  inp.addEventListener("input", function(e) {
      var a, b, i, val = this.value;
      /*close any already open lists of autocompleted values*/
      closeAllLists(this);
      if (!val) { return false;}
      if (val.length <= 1) { return false;}

      currentFocus = -1;
      queryNames.sendRequest(val, this);

  });

  /*execute a function presses a key on the keyboard:*/
  inp.addEventListener("keydown", function(e) {
      var x = document.getElementById(this.id + "autocomplete-list");
      if (x) x = x.getElementsByTagName("div");
      if (e.keyCode == 40) {
        /*If the arrow DOWN key is pressed,
        increase the currentFocus variable:*/
        currentFocus++;
        /*and and make the current item more visible:*/
        addActive(x);
      } else if (e.keyCode == 38) { //up
        /*If the arrow UP key is pressed,
        decrease the currentFocus variable:*/
        currentFocus--; /*and and make the current item more visible:*/
        addActive(x);
      } else if (e.keyCode == 13) {
        /*If the ENTER key is pressed, prevent the form from being submitted,*/
        e.preventDefault();
        if (currentFocus > -1) {
          /*and simulate a click on the "active" item:*/
          if (x) x[currentFocus].click();
        }
      }
  });

  function addActive(x) {
    /*a function to classify an item as "active":*/
    if (!x) return false;
    /*start by removing the "active" class on all items:*/
    removeActive(x);
    if (currentFocus >= x.length) currentFocus = 0;
    if (currentFocus < 0) currentFocus = (x.length - 1);
    /*add class "autocomplete-active":*/
    x[currentFocus].classList.add("autocomplete-active");
  }
  function removeActive(x) {
    /*a function to remove the "active" class from all autocomplete items:*/
    for (var i = 0; i < x.length; i++) {
      x[i].classList.remove("autocomplete-active");
    }
  }
  /*execute a function when someone clicks in the document:*/
  document.addEventListener("click", function (e) {
      closeAllLists(e.target);
  });
}

function closeAllLists(elmnt) {
  /*close all autocomplete lists in the document,
  except the one passed as an argument:*/
  var x = document.getElementsByClassName("autocomplete-items");
  for (var i = 0; i < x.length; i++) {
    if (elmnt != x[i]) {
      x[i].parentNode.removeChild(x[i]);
    }
  }
}

/*initiate the autocomplete function on the "myInput" element, and pass along the countries array as possible autocomplete values:*/
autocomplete(document.getElementById("myInput"));

// for one time non-repeating requests
// irrespective of if it fails or succeeds
// rand_int is used to map requests to response correctly
// every new request resets rand_int,
// effectively making any previous request garbage.
var queryNames = {
  rand_int : 0,
  string: "",
  url: "/query/name/",
  inpel: undefined, // the input element

  callBack: function (data, success) {
    if (success) {
      var elementId = "data-test";
      var obj = JSON.parse(data);
      if (this.rand_int == obj.rand_int) {
        // race condition possible here.
        this.rand_int = 0; // reset
        // TODO : the action to be taken
      //  document.getElementById(elementId).innerHTML = obj.data;

      var a, b, i;
      var arr = obj.data;
      currentFocus = -1;
      /*create a DIV element that will contain the items (values):*/
      a = document.createElement("DIV");
      a.setAttribute("id", this.inpel.id + "autocomplete-list");
      a.setAttribute("class", "autocomplete-items");
      /*append the DIV element as a child of the autocomplete container:*/
      this.inpel.parentNode.appendChild(a);
      var inpel = this.inpel;

      /*for each item in the array...*/
      for (i = 0; i < arr.length; i++) {
        /*check if the item starts with the same letters as the text field value:*/
        /*create a DIV element for each matching element:*/
        b = document.createElement("DIV");
        /*make the matching letters bold:*/
        b.innerHTML = "<strong>" + arr[i][1] + "</strong>";
        //b.innerHTML += arr[i][1];
        /*insert a input field that will hold the current array item's value:*/
        b.innerHTML += "<input type='hidden' value='" + arr[i][1] + "'>";
        /*execute a function when someone clicks on the item value (DIV element):*/
        b.addEventListener("click", function(e) {
            /*insert the value for the autocomplete text field:*/
            inpel.value = this.getElementsByTagName("input")[0].value;
            /*close the list of autocompleted values,
            (or any other open lists of autocompleted values:*/
            closeAllLists(inpel);
        });
        a.appendChild(b);
      }
    } else {
      this.rand_int = 0;
    }
    }
  },

  sendRequest: function (string, inpel) {
    this.string = string;
    this.inpel = inpel;
    // new random variable for new request
    this.rand_int = Math.floor(Math.random() * 1000000000) + 1;
    var query_url = this.url + this.string + "/" + this.rand_int;
    httpGetAsyncObj(query_url, this.callBack, this);
  }
};


</script>

</body>
</html>
