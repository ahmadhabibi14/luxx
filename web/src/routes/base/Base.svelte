<script>
	import { Router, Route, navigate } from "svelte-routing";
	import Home from "./Home.svelte";
	import Navbar from "../../lib/partials/navbar.svelte";
	import Head from "../../lib/partials/head.svelte";
	import Header from "../../lib/partials/header.svelte";
	import Profile from "./Profile.svelte";
	import NotFound from "../error/NotFound.svelte";
	import Unavailable from "../error/Unavailable.svelte";
	import Growl from "../../lib/components/growl.svelte";
	import { getCookie } from "../../lib/utils/helper";
	import { onMount } from "svelte";

	export let url = "/";
	let user_data;
	let growlComponent;
	let onLoading = false;
	onMount(async () => {
		const auth = getCookie("auth");
		if (auth) {
			onLoading = true;
			const resp = await fetch("/api/user/user-data", {
				method: "POST",
				credentials: "same-origin"
			});
			if (resp.ok) {
				onLoading = false;
				const respData = await resp.json();
				user_data = await JSON.parse(JSON.stringify(respData));
			} else if (!resp.ok) {
				onLoading = false;
				const errorData = await resp.json();
				const errorMsg = await JSON.parse(JSON.stringify(errorData));
				// document.cookie = "auth= ; expires = Thu, 01 Jan 1970 00:00:00 GMT";
				growlComponent.showGrowl("error", errorMsg["error"]);
				setTimeout(() => {
					navigate("/login");
				}, 2000);
			}
		} else {
			navigate("/login");
		}
	});
</script>

<Head title="LuXX - Home" description="LuXX" />
<Growl bind:this={growlComponent} />

<Router {url}>
	<div class="w-full bg-zinc-950 h-screen">
		<div class="lg:w-8/12 md:w-10/12 h-full mx-auto flex flex-row gap-3 py-3">
			<Navbar />
			<section class="h-full w-full flex flex-col gap-3">
				<Header {user_data} />
				<div class="route_container">
					{#if onLoading}
						<div class="flex flex-col justify-center items-center gap-3 h-full w-full">
							<svg viewBox="0 0 24 24" class="fill-current animate-spin w-8 h-auto">
								<path
									d="M15.2398 22.75H8.75982C6.86982 22.75 5.48982 21.96 4.95982 20.6C4.40982 19.18 4.91982 17.42 6.23982 16.23L10.8798 12L6.23982 7.77C4.91982 6.58 4.40982 4.82 4.95982 3.4C5.48982 2.03 6.86982 1.25 8.75982 1.25H15.2398C17.1298 1.25 18.5098 2.04 19.0398 3.4C19.5898 4.82 19.0798 6.58 17.7598 7.77L13.1198 12L17.7698 16.23C19.0798 17.42 19.5998 19.18 19.0498 20.6C18.5098 21.96 17.1298 22.75 15.2398 22.75ZM11.9998 13.01L7.24982 17.33C6.40982 18.1 6.03982 19.22 6.35982 20.05C6.65982 20.82 7.50982 21.25 8.75982 21.25H15.2398C16.4898 21.25 17.3398 20.83 17.6398 20.05C17.9598 19.22 17.5998 18.1 16.7498 17.33L11.9998 13.01ZM8.75982 2.75C7.50982 2.75 6.65982 3.17 6.35982 3.95C6.03982 4.78 6.39982 5.9 7.24982 6.67L11.9998 10.99L16.7498 6.67C17.5898 5.9 17.9598 4.78 17.6398 3.95C17.3398 3.18 16.4898 2.75 15.2398 2.75H8.75982Z"
								/>
							</svg>
							<span class="text-sm font-semibold">Loading...</span>
						</div>
					{:else}
						<Route path="/">
							<Home {user_data} />
						</Route>
						<Route path="/profile">
							<Profile />
						</Route>
						<Route path="/chats">
							<Unavailable />
						</Route>
						<Route path="/favorites">
							<Unavailable />
						</Route>
						<Route path="/explore">
							<Unavailable />
						</Route>
						<Route path="/settings">
							<Unavailable />
						</Route>
						<Route component={NotFound} />
					{/if}
				</div>
			</section>
		</div>
	</div>
</Router>

<style lang="postcss">
	.route_container {
		@apply bg-zinc-900 border border-zinc-700 rounded-2xl shadow-sm text-zinc-100 w-full h-full p-4 overflow-y-scroll;
	}
	.route_container::-webkit-scrollbar {
		@apply w-2;
	}
	.route_container::-webkit-scrollbar-thumb {
		@apply rounded-xl bg-zinc-700;
	}
	.route_container::-webkit-scrollbar-track {
		@apply bg-zinc-900 my-3 border-4 border-transparent rounded-full w-2;
	}
</style>
