<script>
	import Navbar from "$lib/partials/navbar.svelte";
	import "../app.css";
</script>

<div class="app">
	<Navbar />

	<main>
		<slot />
	</main>

	<footer>
		<p>visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to learn SvelteKit</p>
	</footer>
</div>