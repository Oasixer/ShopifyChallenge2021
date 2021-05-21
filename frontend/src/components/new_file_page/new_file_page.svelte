<script>
	import TopbarDash from '../general/menu/topbar_dash/topbar.svelte';
	import TagCreate from '../tag_create/tag_create2.svelte';
	import { fileUpload } from '../../store/file.js';
	import { mutation } from '@urql/svelte';

	let title;

	let needTitle = false;

	const fileInfo = JSON.parse($fileUpload);
	const fileExt = fileInfo.fileExt;
	const fileUuid = fileInfo.uuid;

	const getImage = async () => {
		let url = new URL(`${window.location.protocol}//${window.location.host}/api/download-file`);
		url.search = new URLSearchParams({'uuid': fileUuid, 'fileExt': fileExt}).toString();
		const result = await fetch(url);
		return result.blob();
	}

	getImage().then((blob)=>{
    var img = URL.createObjectURL(blob);
    document.getElementById('img1').setAttribute('src', img);
	});

	// $: console.log(JSON.stringify(imgData));
	
	let tags;

	const saveFileMutation = mutation({
			query: `
			mutation($name: String!, $uuid: String!, $tags: String!, $fileExt: String!){
				saveFile(
					name: $name,
					tags: $tags,
					uuid: $uuid,
 					fileExt: $fileExt
				){
					uuid
				}
			}
		`,
	});
	const submitFile = () => {
		if (!title){
			needTitle=true;
			return;
		}
		console.log(`FUCKIN TAGS: ${JSON.stringify(tags)}`);
		console.log(`filtered!!!: ${tags.filter(x=>x.name.length>0).map(x=>x.name).join(',')}`);
		saveFileMutation({
			name: title,
			uuid: fileUuid,
			tags: tags.filter(x=>x.name.length>0).map(x=>x.name).join(','),
			fileExt: fileExt
		}).then(result => {
			if (result.data) {
				console.log('GOT RESULT!!');
				console.log(result.data.saveFile.uuid);
				console.log(JSON.stringify(result.data));
				window.open("/dashboard", "_self");
			} else if (result.error) {
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
			<input type='text' class='title' class:needTitle placeholder='Add title' bind:value={title} on:focus={()=>{needTitle=false}}>
		</div>
		<div class='imgContain'>
			<img id='img1' alt='hi' src=''>
		</div>
		<div class='addTagsLabel'>Add tags</div>
		<TagCreate bind:tags/>
		<button class='submit' on:click={submitFile}>Upload</button>
	</div>
</div>
