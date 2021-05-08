<script>
	import LabeledTextInput from '../labeled_text_input/labeled_text_input.svelte';
	import { email } from '../../../../store/register.js';
  const { open } = getContext('popup-modal');
	import InfoPopup from '../../info_popup/info_popup.svelte'
	import { mutation } from '@urql/svelte';
  import { getContext } from 'svelte';

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
				showCreatedPopup();
			} else if (result.error){
				displayErrorPopup = true;
				console.error(result.error);
			}
		});
	};

	const showCreatedPopup = () => {
		email.set(getInputVal('Email'));
		open(InfoPopup, {message: '', title: 'Account created succesfully!', bottomLink: {'title': 'Login to ImgHub', 'href': '/login'}});
	}

</script>
<style src='create_account_fields.scss'>
</style>

<form id='registerForm' bind:this={form} on:submit|preventDefault={submit}>
	<div id=labeledInputsBlock>
		{#each inputs as input}
			<LabeledTextInput label={input.name} bind:value={input.val} bind:form={form} type={input.name==='Password'?'password':'text'}/>
		{/each}
		<button type='submit' value='Submit' id='registerButton'>
			Register
		</button>
	</div>
</form>
