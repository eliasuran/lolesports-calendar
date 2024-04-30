export interface Result {
	error?: Error;
	data?: any;
}

const apiURL = 'http://localhost:8080/';

export async function mpuFetch(endpoint: string, method: 'GET' | 'POST') {
	try {
		const res = await fetch(apiURL + endpoint, {
			method,
			headers: {
				'Content-type': 'application/json'
			}
		});

		const data = await res.json();

		if (res.status !== 200) {
			console.log(res);
			return { error: data.error };
		}

		return { data };
	} catch (error) {
		return error;
	}
}
