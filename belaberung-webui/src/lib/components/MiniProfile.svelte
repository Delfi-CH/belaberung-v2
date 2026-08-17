<script lang="ts">
	import { getUserDetails } from "$lib/api/core";
	import { onMount } from "svelte";
    import { Card, CardBody, CardHeader, CardTitle } from "@sveltestrap/sveltestrap"

    let { userID } = $props()
    let userData = $state({})
    let biography = $state([])

    onMount(async ()=> {
        const tmpUserData = await getUserDetails(userID)
        userData = tmpUserData
        biography = tmpUserData.biography.split("\n")
    })
</script>

<div class="on-top">
<Card>
    <CardHeader>
        <p>{userData.username}@{userData.domain}</p>
        <p class="pronouns">{userData.pronouns}</p>
    </CardHeader>
    <CardBody>
        {#each biography as line, index (index)}
            <span>{line}</span>
            <br>
        {/each}
    </CardBody>
</Card>
</div>

<style>
    .on-top {
        top: 0.01rem;
        position: absolute;
        z-index: 1000;
    }
    .pronouns {
        font-size: 85%;
    }
</style>