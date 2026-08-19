export const ssr = false;
export const prerender = true;

import { setBackendUrl } from 'belaberung-client-libs';

if (import.meta.env.DEV) { setBackendUrl("http://localhost:8081") }