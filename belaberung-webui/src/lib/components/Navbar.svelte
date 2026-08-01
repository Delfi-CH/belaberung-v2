<script lang="ts">
	import {
		Nav,
		Navbar,
		NavbarBrand,
		NavbarToggler,
		Collapse,
	} from '@sveltestrap/sveltestrap';
	import { page } from '$app/state';
	import { getUsername, isLoggedIn } from '$lib/api/auth';
	import NavbarInsides from './NavbarInsides.svelte';

	let isUserLoggedIn = $state(false);
	let isMobile = $state(false);
	let isToggled = $state(false);
	let username = $state('');

	async function updateLogin() {
		isUserLoggedIn = await isLoggedIn();
	}

	$effect(() => {
		page.route.id;
		isMobile = window.innerWidth < 768;

		void updateLogin();

		username = getUsername();
	});
</script>

<Navbar expand="md">
	<NavbarBrand>belaberung</NavbarBrand>

	{#if isMobile}
		<NavbarToggler onclick={() => (isToggled = !isToggled)} />
		<!-- 
			Yes, this has to be an if statement instead of isOpen={isToggled}. 
			No, i dont know why.
		-->
		{#if isToggled}
		<Collapse isOpen={true} navbar>
			<Nav class="text-end" navbar>
				<NavbarInsides page={page} username={username} isUserLoggedIn={isUserLoggedIn}></NavbarInsides>
			</Nav>
		</Collapse>
		{:else}
		<Collapse isOpen={false} navbar>
			<Nav class="text-end" navbar>
				<NavbarInsides page={page} username={username} isUserLoggedIn={isUserLoggedIn}></NavbarInsides>
			</Nav>
		</Collapse>
		{/if}
		{:else}
		<Nav pills>
			<NavbarInsides page={page} username={username} isUserLoggedIn={isUserLoggedIn}></NavbarInsides>
		</Nav>
	{/if}
</Navbar>
