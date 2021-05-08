<script>
	import LabeledTextInput from '../labeled_text_input/labeled_text_input.svelte';
	import {email} from '../../../../store/register.js';
	let inputs = [{'name':'Username', 'val': ''},{'name':'Email', 'val': `${$email}`},{'name':'Password', 'val': ''}];
	email.set(''); // null out the store after taking the value so that if they reload the page they start over
	let form;

	const getInputVal = (inputName) => {
		for (const i of inputs) {
			if (i.name === inputName){
				return i.val;
			}
		};
	}

	import { mutation } from '@urql/svelte';


	const signupMutation = mutation({
			query: `
			mutation($email: String!, $password: String!, $username: String!){
				signUp(
					email: $email,
					password: $password,
					username: $username
				){
					id,
					email,
					createdAt,
					updatedAt
				}
			}
		`,
	});

	// const signUp = mutation(signupMutation);
	const submit = () => {
		signupMutation({
			email: getInputVal('Email'),
			password: getInputVal('Password'),
			username: getInputVal('Username'),
		}).then(result => {
			if(result.data){
				console.log(result.data);
			} else if (result.error){
				console.error(result.error);
			}
		});
	};

</script>
<style src='create_account_fields.scss'>
</style>

<form id='registerForm' bind:this={form} on:submit|preventDefault={submit}>
	<div id=labeledInputsBlock>
		{#each inputs as input}
			<LabeledTextInput label={input.name} bind:value={input.val} bind:form={form}/>
		{/each}
		<button type='submit' value='Submit' id='registerButton'>
			Register
		</button>
	</div>
</form>
