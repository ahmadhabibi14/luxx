<script>
	import { Router, Route } from "svelte-routing";
	import Home from "./Home.svelte";
	import Navbar from "../../lib/partials/navbar.svelte";
	import Head from "../../lib/partials/head.svelte";
	import Header from "../../lib/partials/header.svelte";
	import Profile from "./Profile.svelte";

	export let url = "/";
</script>

<Head title="LuXX - Home" description="LuXX" />

<Router {url}>
	<div class="w-full bg-slate-200 h-screen">
		<div class="md:w-8/12 h-full mx-auto flex flex-row gap-6 py-6">
			<Navbar />
			<section class="h-full w-full flex flex-col gap-6">
				<Header />
				<div
					class="bg-slate-50 border border-slate-300 rounded-2xl shadow-sm text-slate-800 w-full h-full py-6 px-4"
				>
					<Route path="/" component={Home} />
					<Route path="/profile" component={Profile} />
				</div>
			</section>
		</div>
	</div>
</Router>
