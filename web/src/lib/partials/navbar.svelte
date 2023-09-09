<script>
	import { Link } from "svelte-routing";
</script>

<div class="navbar">
	<div class="navbarDesktop">
		<nav class="w-fit flex flex-row gap-14 justify-center content-center items-center">
			<Link to="/">Home</Link>
			<Link to="/about">About</Link>
			<Link to="/blog">Blog</Link>
		</nav>
	</div>
</div>
