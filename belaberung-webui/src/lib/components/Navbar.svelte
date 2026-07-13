<script lang="ts">
	import {
		Nav,
		Navbar,
		NavbarBrand,
		NavbarToggler,
		NavItem,
		NavLink,
        Dropdown,
        DropdownItem,
        DropdownMenu,
        DropdownToggle,
        Icon
	} from '@sveltestrap/sveltestrap';
    import { page } from '$app/state';
    import { getUsername, isLoggedIn } from '$lib/api/auth';

    let isUserLoggedIn = $state(false);
    let username = $state("")

    async function updateLogin() {
        isUserLoggedIn = await isLoggedIn();
    }

    $effect(() => {
        page.route.id;

        void updateLogin();

        username = getUsername()
    });
</script>

<Navbar>
	<NavbarBrand>belaberung</NavbarBrand>
	<Nav pills>	
		{#if isUserLoggedIn}
            <NavItem>
			    <NavLink href="/" active={page.route.id == "/"}>Home</NavLink>
		    </NavItem>
			<NavItem>
				<NavLink href="/rooms" active={page.route.id == "/rooms"}>Rooms</NavLink>
			</NavItem>
            <Dropdown autoClose="manual" isOpen={false}>
                <DropdownToggle caret><Icon name="person-circle"/> {username}
                </DropdownToggle>
                <DropdownMenu>
                    <DropdownItem>Your Profile</DropdownItem>
                    <DropdownItem href="/logout" class="bg-danger text-white">Logout</DropdownItem>
                </DropdownMenu>
            </Dropdown>
		{:else}
			<NavItem>
				<NavLink href="/login"  active={page.route.id == "/login"}>Login</NavLink>
			</NavItem>
			<NavItem>
				<NavLink href="/register"  active={page.route.id == "/register"}>Register</NavLink>
			</NavItem>
		{/if}
	</Nav>
</Navbar>
