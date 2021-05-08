<script>
	import { query, operationStore } from "@urql/svelte";
	const getUser = operationStore(`
    query {
			getUser {
				files {
					documentType
				}
			}
		}
  `);
		query(getUser);
		getUser.subscribe(value => {
			if(value.fetching){
					console.log("loading user files");
			}
			else if(value.error){
					console.error("graphql error loading user: " + value.error)
			}
			else{
				console.log(`graphql value (user): ${value.data.getUser.files}`);
			}
		});
</script>
<style src='dashboard.scss'>
</style>
