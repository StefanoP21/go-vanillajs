import { HomePage } from "./components/HomePage.js";

window.addEventListener("DOMContentLoaded", (event) => {
	document.querySelector("main").appendChild(new HomePage());
});

window.app = {
	search: (event) => {
		event.preventDefault();
		const keywords = document.querySelector("input[type=search]").value;
	},
};

window.addEventListener("DOMContentLoaded", () => {});
