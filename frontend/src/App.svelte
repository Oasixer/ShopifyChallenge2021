<script>
	import { Router } from "@roxi/routify";
	import { routes } from "../.routify/routes"; // this is a build time import
	import { sessionToken } from "./store/session.js"; // this is a build time import
	
	// graphql stuff
	import { initClient, query, operationStore, dedupExchange, cacheExchange, fetchExchange } from "@urql/svelte";
	import { authExchange } from "@urql/exchange-auth";
	import { makeOperation } from '@urql/core';

	const addAuthToOperation = ({ authState, operation }) => {
		if (!authState || !authState.token) {
			return operation;
		}

		const fetchOptions =
			typeof operation.context.fetchOptions === 'function'
				? operation.context.fetchOptions()
				: operation.context.fetchOptions || {};

		return makeOperation(operation.kind, operation, {
			...operation.context,
			fetchOptions: {
				...fetchOptions,
				headers: {
					...fetchOptions.headers,
					Authorization: authState.token,
				},
			},
		});
	};

	const url = `${window.location.protocol}//${window.location.host}/graphql`;

	const logout = () => {
		sessionToken.set('');
		console.log('logging out');
		// window.open('/login');
	}

	const getAuth = async ({ authState }) => {
		if (!authState) {
			const token = $sessionToken
			// const refreshToken = localStorage.getItem('refreshToken');
			// if (token && refreshToken) {
				// return { token, refreshToken };
			// }
			if (token){
				return { token };
			}
			return null;
		}

		logout();

		return null;
	};

	initClient({
		url: url,
		exchanges: [
			dedupExchange,
			cacheExchange,
			authExchange({
				addAuthToOperation: addAuthToOperation,
				getAuth: getAuth
			}),
			fetchExchange
		]
	});

	const test = operationStore(`
	 {
			test
	 }
	`);

	query(test)

	test.subscribe(value => {
		if(value.fetching){
				console.log("loading");
		}
		else if(value.error){
				console.error("graphql error: " + value.error)
		}
		else{
			console.log(`graphql value: ${value.data.test}`);
		}
	});
</script>

<style src="./app.scss" global></style>

<Router {routes}/>
