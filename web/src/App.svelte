<script>
	import { onMount } from "svelte";
	import { Route, Router, navigate } from "svelte-routing";
	import "./app.css";
	import Base from "./routes/base/Base.svelte";
	import NotFound from "./routes/error/NotFound.svelte";
	import Register from "./routes/auth/Register.svelte";
	import Login from "./routes/auth/Login.svelte";
	import { getCookie } from "./lib/utils/helper";

	onMount(() => {
		const auth = getCookie("auth");
		if (auth) {
			console.log(auth);
			navigate("/");
		} else {
			navigate("/login");
		}
	});
</script>

<Router>
	<Route path="/*" component={Base} />
	<Route path="/register" component={Register} />
	<Route path="/login" component={Login} />
	<Route component={NotFound} />
</Router>
