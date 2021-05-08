<script>
	import LabeledTextInput from '../labeled_text_input/labeled_text_input.svelte';
	let inputs = [{'name':'Username', 'val': ''},{'name':'Password', 'val': ''}];
	const submit = () => {};
	let form;
</script>
<style src='./login_fields.scss'>
</style>

<form id='loginForm' bind:this={form}>
	<div id=labeledInputsBlock>
		{#each inputs as input}
			<LabeledTextInput label={input.name} bind:value={input.val} bind:form={form}/>
		{/each}
		<button type='submit' value='Submit' id='loginButton' on:click={submit}>
			Login
		</button>
	</div>
</form>
