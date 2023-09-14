<script>
	import { Router, Route } from "svelte-routing";
	import Home from "./Home.svelte";
	import Navbar from "../../lib/partials/navbar.svelte";

	export let url = "";
</script>

<Router {url}>
	<Navbar />
	<div class="bg-slate-950">
		<Route path="/">
			<Home />
		</Route>
	</div>
</Router>
