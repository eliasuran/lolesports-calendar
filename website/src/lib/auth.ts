import { mpuFetch, type Result } from './lib.js';

export async function authURL() {
	const res = (await mpuFetch('auth', 'GET')) as Result;
	if (res.error) {
		console.log(res.error);
		return res.error;
	}
	return res.data;
}
