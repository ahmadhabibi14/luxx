<script>
	import { Router, Link, Route } from "svelte-routing";
	import Home from "./Home.svelte";

	export let url = "";
</script>

<Router {url}>
	<nav>
		<Link to="/">Home</Link>
		<Link to="/about">About</Link>
		<Link to="/blog">Blog</Link>
	</nav>
	<div>
		<Route path="/"><Home /></Route>
	</div>
</Router>
