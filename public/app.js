import "./components/AnimateLoading.js";
import { Router } from "./services/Router.js";
import { API } from "./services/API.js";

window.app = {
	Router,
	search: (event) => {
		event.preventDefault();
		const keywords = document.querySelector("input[type=search]").value;
	},
	api: API,
};

window.addEventListener("DOMContentLoaded", () => {
	app.Router.init();
});
