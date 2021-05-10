<script>
	import TopbarDash from '../general/menu/topbar_dash/topbar.svelte';
	import TagCreate from '../tag_create/tag_create.svelte';
	import { fileUpload } from '../../store/file.js';
	import { mutation } from '@urql/svelte';

	// const base64 = require('base-64');

	// function _arrayBufferToBase64( buffer ) {
		// var binary = '';
		// var bytes = new Uint8Array( buffer );
		// var len = bytes.byteLength;
		// for (var i = 0; i < len; i++) {
				// binary += String.fromCharCode( bytes[ i ] );
		// }
		// return base64.encode( binary );
	// }
	// let files = ['https://images.unsplash.com/photo-1481349518771-20055b2a7b24?ixid=MnwxMjA3fDB8MHxzZWFyY2h8M3x8cmFuZG9tfGVufDB8fDB8fA%3D%3D&ixlib=rb-1.2.1&w=1000&q=80'];

	let title;

	const fileInfo = JSON.parse($fileUpload);
	console.log(`fileInfo: ${JSON.stringify(fileInfo)}`);
	const fileExt = fileInfo.fileExt;
	const fileUuid = fileInfo.uuid;
	console.log(`fileExt: ${fileInfo.fileExt}`);

	const getImage = async () => {
		let url = new URL(`${window.location.protocol}//${window.location.host}/api/download-file`);
		url.search = new URLSearchParams({'uuid': fileUuid, 'fileExt': fileExt}).toString();
		const result = await fetch(url);
		return result.blob();
	}



	// function hexToBase64(str) {
		// return btoa(String.fromCharCode.apply(null, str.replace(/\r|\n/g, "").replace(/([\da-fA-F]{2}) ?/g, "0x$1 ").replace(/ +$/, "").split(" ")));
	// }
	let fr = new FileReader();
	let imgData;
	fr.onload = function ( oFREvent ) {
		imgData = btoa(oFREvent.target.result);
	}
	getImage().then((blob)=>{
		// console.log(`fetched.${JSON.stringify(blob)}`);
		// fr.readAsText(blob, "utf-8")
		// console.log(`fetched.${imgData}`);
    var img = URL.createObjectURL(blob);
    // Do whatever with the img
    document.getElementById('img1').setAttribute('src', img);
	});

	$: console.log(JSON.stringify(imgData));
	
	let tags;

	const saveFileMutation = mutation({
			query: `
			mutation($name: String!, $uuid: String!, $tags: String!){
				saveFile(
					name: $name,
					tags: $tags,
					uuid: $uuid
				){
					uuid
				}
			}
		`,
	});
	const submitFile = () => {
		saveFileMutation({
			name: title,
			uuid: fileUuid,
			tags: tags.filter(x=>x.name.length>0).map(x=>x.name).join(',')
		}).then(result => {
			if (result.data) {
				console.log('GOT RESULT!!');
				console.log(result.data.uuid);
				console.log(JSON.stringify(result.data));
				// window.open();
			} else if (result.error) {
				// displayErrorPopup = true;
				console.error(result.error);
			}
		});
	}

</script>
<style src='new_file_page.scss'>
</style>
<div id='newFilePageMain'>
	<TopbarDash showUpload={false}/>
	<div class='content'>
		<div class='titleContain'>
			<input type='text' class='title' placeholder='Add title' bind:value={title}>
		</div>
		<div class='imgContain'>
			<img id='img1' alt='hi' src=''>
		</div>
		<TagCreate bind:tags/>
		<button class='submit' on:click={submitFile}>Upload</button>
	</div>
</div>
