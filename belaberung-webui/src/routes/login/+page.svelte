<script lang="ts">
	import { api } from "$lib/api";
	import axios from "axios";

    let username = $state("")
    let password = $state("")
    let errorMessage = $state("")

    async function handleSubmit() {
        try {
            const res = await api.post("/auth/login", {
            username: username,
            password: password
            })
            console.log(res.data)
        } catch (err) {
            console.error(err)
            if (axios.isAxiosError(err)) {
                if (err.status === 404) {
                    errorMessage = "user doesnt exist"
                } else if (err.status === 418) {
                    errorMessage = "already logged in"
                } else if (err.status === 400) {
                    errorMessage = "bad request"
                } else if (err.status === 500) {
                    errorMessage = "server error"
                } else if (err.status === 401) {
                    errorMessage = "wrong password"
                } else {
                    errorMessage = "unexpected error: " + err
                }
            }
        }
    }
</script>

<form onsubmit={async()=> await handleSubmit()}>
    <label for="username">Username</label><input type="username" id="username" required bind:value={username}>
    <label for="password">Password</label><input type="password" id="password" required bind:value={password}>
    <button type="submit">Login</button>
    <p>{errorMessage}</p>
</form>