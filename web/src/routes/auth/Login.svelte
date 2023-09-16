<script>
	import { onMount } from "svelte";
	import { navigate } from "svelte-routing";
	import { getCookie } from "../../lib/utils/helper";
	import Head from "../../lib/partials/head.svelte";
	import Growl from "../../lib/components/growl.svelte";

	let growlComponent;
	let onUploading = false;
	onMount(() => {
		const auth = getCookie("auth");
		if (auth) {
			navigate("/");
		}
	});

	let errorMessage = "";
	let email, password;

	async function handleLogin() {
		onUploading = true;
		try {
			const resp = await fetch("/api/auth/login", {
				method: "POST",
				headers: {
					"Content-Type": "application/json"
				},
				body: JSON.stringify({
					email: email,
					password: password
				})
			});

			if (resp.ok) {
				onUploading = false;
				const creds = await resp.json();
				navigate("/");
			} else {
				onUploading = false;
				const errorData = await resp.json();
				const errorMsg = await JSON.parse(errorData);
				growlComponent.showGrowl("error", errorMsg["error"]);
			}
		} catch (error) {
			onUploading = false;
			growlComponent.showGrowl("error", "Unexpected error occurred during register");
		}
	}
</script>

<Head title="Login" description="LuXX - Register" />
<Growl bind:this={growlComponent} />

<main class="bg-slate-100 h-screen w-full">
	<div
		class="px-5 md:px-10 flex flex-row justify-center items-center h-full w-full gap-28 text-slate-950"
	>
		<div class="md:block hidden">
			<img src="/icons/master/laptop.svg" alt="" class="w-[420px] h-auto" />
		</div>
		<div class="w-[450px] bg-white p-6 rounded-xl drop-shadow-lg flex flex-col gap-3">
			<header>
				<h1 class="font-bold text-center text-4xl">Login 🫰</h1>
				<p class="text-center text-sm text-slate-700 mt-1.5">Please enter your details</p>
			</header>
			{#if errorMessage}
				<div
					class="w-full h-fit p-2 rounded-md flex flex-row justify-center gap-2 items-center bg-red-500 text-white"
				>
					<svg viewBox="0 0 24 24" class="w-[18px] h-auto fill-current">
						<path
							d="M12 22.75C6.07 22.75 1.25 17.93 1.25 12C1.25 6.07 6.07 1.25 12 1.25C17.93 1.25 22.75 6.07 22.75 12C22.75 17.93 17.93 22.75 12 22.75ZM12 2.75C6.9 2.75 2.75 6.9 2.75 12C2.75 17.1 6.9 21.25 12 21.25C17.1 21.25 21.25 17.1 21.25 12C21.25 6.9 17.1 2.75 12 2.75Z"
						/>
						<path
							d="M4.89984 19.7499C4.70984 19.7499 4.51984 19.6799 4.36984 19.5299C4.07984 19.2399 4.07984 18.7599 4.36984 18.4699L18.3698 4.46994C18.6598 4.17994 19.1398 4.17994 19.4298 4.46994C19.7198 4.75994 19.7198 5.23994 19.4298 5.52994L5.42984 19.5299C5.27984 19.6799 5.08984 19.7499 4.89984 19.7499Z"
						/>
					</svg>
					<span>{errorMessage}</span>
				</div>
			{/if}
			<div class="flex flex-col gap-4">
				<label for="email">
					<input
						bind:value={email}
						class="w-full border-slate-400/80 border rounded-md bg-transparent focus:border-slate-950 caret-slate-950 py-2 px-4 outline-1 focus:outline-slate-950"
						type="email"
						id="email"
						placeholder="Your Email"
					/>
				</label>
				<label for="password">
					<input
						bind:value={password}
						class="w-full border-slate-400/80 border rounded-md bg-transparent focus:border-slate-950 caret-slate-950 py-2 px-4 outline-1 focus:outline-slate-950"
						type="password"
						id="password"
						placeholder="Password"
					/>
				</label>
				<button
					on:click={handleLogin}
					class="flex justify-center cursor-pointer py-2.5 bg-slate-950 hover:bg-slate-900 hover:shadow-md text-slate-100 rounded-md"
				>
					{#if onUploading}
						<svg width="24" height="24" viewBox="0 0 24 24" class="fill-current animate-spin">
							<path
								d="M15.2398 22.75H8.75982C6.86982 22.75 5.48982 21.96 4.95982 20.6C4.40982 19.18 4.91982 17.42 6.23982 16.23L10.8798 12L6.23982 7.77C4.91982 6.58 4.40982 4.82 4.95982 3.4C5.48982 2.03 6.86982 1.25 8.75982 1.25H15.2398C17.1298 1.25 18.5098 2.04 19.0398 3.4C19.5898 4.82 19.0798 6.58 17.7598 7.77L13.1198 12L17.7698 16.23C19.0798 17.42 19.5998 19.18 19.0498 20.6C18.5098 21.96 17.1298 22.75 15.2398 22.75ZM11.9998 13.01L7.24982 17.33C6.40982 18.1 6.03982 19.22 6.35982 20.05C6.65982 20.82 7.50982 21.25 8.75982 21.25H15.2398C16.4898 21.25 17.3398 20.83 17.6398 20.05C17.9598 19.22 17.5998 18.1 16.7498 17.33L11.9998 13.01ZM8.75982 2.75C7.50982 2.75 6.65982 3.17 6.35982 3.95C6.03982 4.78 6.39982 5.9 7.24982 6.67L11.9998 10.99L16.7498 6.67C17.5898 5.9 17.9598 4.78 17.6398 3.95C17.3398 3.18 16.4898 2.75 15.2398 2.75H8.75982Z"
							/>
						</svg>
					{:else}
						<span>Login</span>
					{/if}
				</button>
				<div class="flex flex-row justify-center items-center gap-2 text-sm">
					<span>Don't have an account?</span>
					<button
						on:click={() => navigate("/register")}
						class="text-sky-500 font-semibold hover:underline">Register</button
					>
				</div>
			</div>
		</div>
	</div>
</main>
