<script lang="ts">
    import {
		NavItem,
		NavLink,
		Dropdown,
		DropdownItem,
		DropdownMenu,
		DropdownToggle,
		Icon
	} from '@sveltestrap/sveltestrap';
    let { isUserLoggedIn, page, username} = $props()
</script>

{#if isUserLoggedIn}
	<NavItem>
		<NavLink href="/" active={page.route.id == '/'}>Home</NavLink>
	</NavItem>
	<NavItem>
		<NavLink href="/rooms" active={page.route.id == '/rooms'}>Rooms</NavLink>
	</NavItem>
	<Dropdown autoClose="manual" isOpen={false}>
		<DropdownToggle caret><Icon name="person-circle" /> {username}</DropdownToggle>
		<DropdownMenu>
            <DropdownItem href="/rooms/new">Create a room</DropdownItem>
			<DropdownItem>Your Profile</DropdownItem>
			<DropdownItem href="/logout" class="bg-danger text-white">Logout</DropdownItem>
		</DropdownMenu>
	</Dropdown>
{:else}
	<NavItem>
		<NavLink href="/login" active={page.route.id == '/login'}>Login</NavLink>
	</NavItem>
	<NavItem>
		<NavLink href="/register" active={page.route.id == '/register'}>Register</NavLink>
	</NavItem>
{/if}
