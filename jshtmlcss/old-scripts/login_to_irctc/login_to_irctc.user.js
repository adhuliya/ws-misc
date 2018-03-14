// ==UserScript==
// @name           login to irctc
// @namespace      neo@gmail.com
// @description    login to irctc automatically
// @include        https://www.irctc.co.in/*
// @include        http://www.irctc.co.in/*
// ==/UserScript==
function login() {
	var form = document.forms.namedItem("LoginForm");
	if(form) {
		alert("going to login automatically");
		form.elements.namedItem("userName").value = "adhuliya";
		form.elements.namedItem("password").value = "anshuisneo";
		form.elements.namedItem("button").click();
	}
}
login();