<script>
	export let label;
	export let value;
	export let form;
	export let type='text';
</script>
<style src='./labeled_text_input.scss'></style>

<div class='labeled-text-input'>
	<div class='label'>{label}</div>
	<!-- its dumb but svelte doesnt want dynamic types w/ two way bindings for some reason.
	 so I cant just do {type} :( -->
	{#if type==='password'}
		<input type='password' name={label.toLowerCase()} bind:value/>
	{:else if type==='text'}
		<input type='text' name={label.toLowerCase()} bind:value/>
	{/if}
</div>
