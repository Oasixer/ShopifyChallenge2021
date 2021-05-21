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
	const uploadFile = async (file) => {
		const formData = new FormData();
		formData.append('file_input', file);
		console.log('running fetch');
		const result = await fetch(`${window.location.protocol}//${window.location.host}/api/upload-file`, {
			method: 'POST',
			body: formData
			});
		// console.log(`Success: ${result.json()}`);
		return result.json();
	}
	const handleFile = (file) => {
		uploadFile(file).then(result => {
			result.fileExt = file.name.split('.').pop();
			console.log(`got upload result: ${JSON.stringify(result)}`);
			fileUpload.set(JSON.stringify(result));
			window.open('/new', '_self');
		});
	}

	let dropArea;

import { onMount } from 'svelte';
	onMount(async => {
		// Prevent default drag behaviors
		;['dragenter', 'dragover', 'dragleave', 'drop'].forEach(eventName => {
			dropArea.addEventListener(eventName, preventDefaults, false)   
			document.body.addEventListener(eventName, preventDefaults, false)
		})

		// Highlight drop area when item is dragged over it
		;['dragenter', 'dragover'].forEach(eventName => {
			dropArea.addEventListener(eventName, highlight, false)
		})

		;['dragleave', 'drop'].forEach(eventName => {
			dropArea.addEventListener(eventName, unhighlight, false)
		})

		// Handle dropped files
		dropArea.addEventListener('drop', handleDrop, false)
	}); 

	function preventDefaults (e) {
		e.preventDefault()
		e.stopPropagation()
	}
	
	let highlighted = false;

	function highlight(e) {
		highlighted = true;
	}

	function unhighlight(e) {
		highlighted = false;
		dropArea.classList.remove('active')
	}

	function handleDrop(e) {
		var dt = e.dataTransfer
		var files = dt.files

		handleFile(files[0]);
	}
</script>
<style src='upload_popup.scss'>
</style>
<input type='file' bind:this={fileInput} style='visibility: hidden; height: 0; width: 0;'
			 on:change={()=>{handleFile(fileInput.files[0])}} accept=".jpg, .jpeg, .png">
<div class=uploadPopup bind:this={dropArea}>
	<div class='dropImage' class:highlighted>
		Drop image here
	</div>
	<div class='icon-wrapper' on:click={()=>fileInput.click()}>
		<Icon data="{imageIcon}" size="50px" viewBox="0 -10 60 80" stroke=""/>
		<div style='margin: 0 10px'></div>
		Choose photo
	</div>
</div>
