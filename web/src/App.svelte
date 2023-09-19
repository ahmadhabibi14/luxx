<script>
	import { Route, Router } from "svelte-routing";
	import "./app.css";
	import Base from "./routes/base/Base.svelte";
	import NotFound from "./routes/error/NotFound.svelte";
	import Register from "./routes/auth/Register.svelte";
	import Login from "./routes/auth/Login.svelte";
</script>

<Router>
	<Route path="/*" component={Base} />
	<Route path="/register" component={Register} />
	<Route path="/login" component={Login} />
</Router>
