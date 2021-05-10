<script>
	import Icon from 'svelte-icon';
	import { fileUpload } from '../../../store/file.js';
	import imageIcon from '../../../../public/assets/icons/image.svg';
	let fileInput;

	fileUpload.set('');

	// import { mutation } from '@urql/svelte';
	// const confirmFileMutation = mutation({
			// query: `
			// mutation($name: String!, $email: String!){
				// signIn(
					// password: $password,
					// email: $email
				// )}`,
	// });

	// const signUp = mutation(signupMutation);
	const uploadFile = async () => {
		const formData = new FormData();
		formData.append('file_input', fileInput.files[0]);
		console.log('running fetch');
		const result = await fetch(`${window.location.protocol}//${window.location.host}/api/upload-file`, {
			method: 'POST',
			body: formData
			});
		// console.log(`Success: ${result.json()}`);
		return result.json();
	}
	const handleFile = () => {
		uploadFile().then(result => {
			result.fileExt = fileInput.files[0].name.split('.').pop();
			console.log(`got upload result: ${JSON.stringify(result)}`);
			fileUpload.set(JSON.stringify(result));
			window.open('/new', '_self');
		});
	}
</script>
<style src='upload_popup.scss'>
</style>
<input type='file' bind:this={fileInput} style='visibility: hidden; height: 0; width: 0;'
	 on:change={handleFile} accept=".jpg, .jpeg, .png">
<div class=uploadPopup>
	<div class='dropImage'>
		Drop image here
	</div>
	<div class='icon-wrapper' on:click={()=>fileInput.click()}>
		<Icon data="{imageIcon}" size="50px" viewBox="0 -10 60 80" stroke=""/>
		<div style='margin: 0 10px'></div>
		Choose photo
	</div>
</div>
