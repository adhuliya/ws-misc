// ==UserScript==
// @name           onlineSBI clicks Confirm to pay automatically
// @namespace      neo@gmail.com
// @description    clicks confirm automatically to pay
// @include        https://www.onlinesbi.com/merchant/merchantdisplay.htm*
// ==/UserScript==
function clickToPay() {
	//alert("hello 11111156");
    var form = document.forms.namedItem('form1');
	var radioRow = document.getElementById("dr0");
	var radioObj = document.getElementById("drd0");
	if(!radioRow) {
		alert("hello no radioRow");
		return;
	}
	if(!radioObj) {
		alert("hello no radioObj");
		return;
	}
	//alert("hello 111114 rsdioID = " + radioRow.id);
	//var radioLength = radioObj.length;
	//alert("hello radioLength = " + radioLength);
	//changeRowColor(radioRow.id,'click','tblAcctd','d');
	//selectAccountNo('00000020008531321','00000020008531321','A1','02623','debitAccountNo','debitAccountType','debitBranchCode','d');
	//setBankSystem('SBCHQ-GEN-PUB-IND-NONRURAL-INR','Core','d');
	//alert(radioRow.onClick);
	radioObj.checked = true;
	radioObj.click();
	//radioRow.click();
	//alert("hello 111119");
    if(form.elements.namedItem('Submit')) {
		var submitButton = form.elements.namedItem('Submit');
			submitButton.click();
    }
}
clickToPay();