import { isLoggedIn } from '$lib/api/auth';
import { error, redirect } from '@sveltejs/kit';
import { api } from "$lib/api/core"

export const ssr = false;
export const prerender = false;


export async function load({ params }) {
	if (await isLoggedIn()) {
		try {
            const res = await api.get("/rooms/" + params.slug)
            return {
                post: res.data
            }
        } catch {
            error(500, "server error")
        }
	} else {
		throw redirect(307, '/login');
	}
}
