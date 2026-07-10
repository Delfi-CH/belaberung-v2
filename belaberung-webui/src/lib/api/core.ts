import axios from 'axios';

const backendURL = import.meta.env.DEV ? 'http://localhost:8081' : '/api';

export const api = axios.create({
	baseURL: backendURL,
	withCredentials: true
});
