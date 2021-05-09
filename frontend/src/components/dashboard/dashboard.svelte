<script>
	import { query, operationStore } from "@urql/svelte";
	import TopbarDash from '../general/menu/topbar_dash/topbar.svelte';
	const getUser = operationStore(`
    query {
			getUser {
				files {
					name
					tags
					url
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

	let files = ['https://images.unsplash.com/photo-1481349518771-20055b2a7b24?ixid=MnwxMjA3fDB8MHxzZWFyY2h8M3x8cmFuZG9tfGVufDB8fDB8fA%3D%3D&ixlib=rb-1.2.1&w=1000&q=80'];
</script>
<style src='dashboard.scss'>
</style>
<div id='dashboardMain'>
	<TopbarDash/>
	{#each files as img}
		<img src={img} alt='thing'>
	{/each}
</div>
