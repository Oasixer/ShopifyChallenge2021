<script>
	import LabeledTextInput from '../labeled_text_input/labeled_text_input.svelte';
	import { email } from '../../../../store/register.js';
	import { sessionToken } from '../../../../store/session.js';
	import { mutation } from '@urql/svelte';
	let inputs = [{'name':'Email', 'val': ''},{'name':'Password', 'val': ''}];
	let form;
	const setInputVal = (inputName, val) => {
		for (const i of inputs) {
			if (i.name === inputName){
					i.val = val;
			}
		};
	}
	if ($email){
		setInputVal('Email', $email);
		email.set(''); // clear cached email
	}

	const getInputVal = (inputName) => {
		for (const i of inputs) {
			if (i.name === inputName){
				return i.val;
			}
		};
	}

	const signinMutation = mutation({
			query: `
			mutation($password: String!, $email: String!){
				signIn(
					password: $password,
					email: $email
				)}`,
	});

	// const signUp = mutation(signupMutation);
	const submit = () => {
		signinMutation({
			password: getInputVal('Password'),
			email: getInputVal('Email'),
		}).then(result => {
			console.log(`result: ${result}`);
			if (result.error){
				console.error(result.error);
			}
			else if (result.data){
				console.log(`result.data.signIn: ${result.data.signIn}`);
				sessionToken.set(result.data.signIn);
				window.open("/dashboard", "_self");
			}
		});
	};
</script>
<style src='./login_fields.scss'>
</style>

<form id='loginForm' bind:this={form} on:submit|preventDefault={submit}>
	<div id=labeledInputsBlock>
		{#each inputs as input}
			<LabeledTextInput label={input.name} bind:value={input.val} type={input.name==='Password'?'password':'text'}/>
		{/each}
		<button type='submit' value='Submit' id='loginButton' on:click={submit}>
			Login
		</button>
	</div>
</form>
