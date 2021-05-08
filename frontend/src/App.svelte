<script>
	import { Router } from "@roxi/routify";
	import { routes } from "../.routify/routes"; // this is a build time import
	
	// graphql stuff
	import { initClient, query, operationStore } from "@urql/svelte";

	console.log(`protocol: ${window.location.protocol}`);
	const url = `${window.location.protocol}//${window.location.host}/graphql`;
	console.log(`url: ${url}`);

	initClient({
			url: url
	})

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
