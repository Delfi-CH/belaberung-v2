<script lang="ts">
	import { api } from "$lib/api";
	import { onMount } from "svelte";

    let errorMessage = $state("")

    onMount(async ()=>{
        try {
            await api.get("/auth/logout")
        } catch {
            errorMessage = "server error"
        }
        
    })
</script>

<p>logout</p>
<p>{errorMessage}</p>