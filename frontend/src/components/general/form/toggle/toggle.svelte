<script>
	export let on;
</script>
<style src='toggle.scss'>
</style>

<label class="switch">
	<input type="checkbox" bind:checked={on}>
	<span class="slider round"></span>
</label>
