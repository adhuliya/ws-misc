// ==UserScript==
// @name           onlineSBi final confirmation
// @namespace      neo@gmail.com
// @description    finally confirm the payment
// @include        https://www.onlinesbi.com/merchant/merchantinter.htm*
// ==/UserScript==
function confirmPayment() {
	var form = document.forms.namedItem('merchantConfirm');
	var confirmButton = document.getElementById('confirmButton');
	if(!confirmButton) {
		return;
	}
	confirmButton.click();
}
confirmPayment();