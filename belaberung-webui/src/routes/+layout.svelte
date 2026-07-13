<script lang="ts">
	import favicon from '$lib/assets/favicon.svg';
	import Navbar from '$lib/components/Navbar.svelte';
	import "bootstrap-icons/font/bootstrap-icons.css";

	let { children } = $props();
	let theme = $state("zephyr")
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
	<link rel="stylesheet" href={`../../node_modules/bootswatch/dist/${theme}/bootstrap.css`}>
</svelte:head>
<Navbar></Navbar>
{@render children()}
