<script>
	import Toggle from '../toggle/toggle.svelte';
	export let items;
</script>
<style src='toggle_list.scss'>
</style>
<div class='toggle_list'>
	{#each items as item}
		<div class='labeled_toggle'>
			<div class='title'>{item.title}</div>
			<Toggle bind:on={item.on}/>
		</div>
	{/each}
</div>
