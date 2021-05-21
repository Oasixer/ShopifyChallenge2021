<script>
	import { query, operationStore } from "@urql/svelte";
	import TopbarDash from '../general/menu/topbar_dash/topbar.svelte';
	import Searchbar from '../general/form/searchbar/searchbar.svelte';

	// const placeholderURL = 'https://images.unsplash.com/photo-1481349518771-20055b2a7b24?ixid=MnwxMjA3fDB8MHxzZWFyY2h8M3x8cmFuZG9tfGVufDB8fDB8fA%3D%3D&ixlib=rb-1.2.1&w=1000&q=80';
	let files = [];
	// const INIT_MAX_FILES = 5; // TODO make this higher
	
	let loaded = false;

	$: hasNoImages = loaded && !files.length;

	let searchtext='';
	let filterText='';

	// let filterText='';

	// $: filterImage = (img) => {
		// if (!filterText){
			// return true;
		// }
		// let tags = img.tags.split(',');
		// if (tags.includes(filterText)){
			// return true;
		// }
		// return false;
	// }
		//
	$: showFile = (file) => {
		if (!filterText){
			return true;
		}
		return numMatches(file.tags) !== 0;
	}
		
	const search = () => {
		console.log('HIASJIDOASJOD');
		filterText = searchtext;
	}

	const getImage = async (file) => {
		console.log('gonna print or wtf');
		let url = new URL(`${window.location.protocol}//${window.location.host}/api/download-file`);
		console.log(`getting image w/ uuid: ${file.uuid}`);
		url.search = new URLSearchParams({'uuid': file.uuid, 'fileExt': file.fileExt}).toString(); // temp fix the fileExt shit
		const result = await fetch(url);
		return result.blob();
	}

	const getBlobs = async (files) => {
		let newFiles = [];
		for (const i of files){
			let blob = await getImage(i);
			i.data = URL.createObjectURL(blob);
			i.tags = i.tags.split(',').map(x=>{return {name: x}});
			newFiles.push(i);
		};
		return newFiles;
	}
	
	const loadFiles = async (graphQLFiles) => {
		files = await getBlobs(graphQLFiles);
		loaded = true;
	}
	
	const getUser = operationStore(`
	query {
			getUser {
				files {
					name
					tags
					uuid
				}
			}
		}
  `);
	query(getUser);
	getUser.subscribe(value => {
		if(value.fetching){
			console.log("loading user files");
		}
		else if(value.error){
			console.error("graphql error loading user: " + value.error)
			console.log(JSON.stringify(value))
			console.log(JSON.stringify(value.error.graphQLErrors))
			value.error.graphQLErrors.forEach(function (i) {
				if (i.extensions.code == "NotFound"){
					window.open('/logout', '_self')
				}
			});
		}
		else{
			console.log(`graphql value (user): ${JSON.stringify(value.data.getUser.files)}`);
			loadFiles(value.data.getUser.files)
		}
	});
	
	const numMatches = (tags) => {
		let num = 0;
		tags.forEach(function (i) {
			if (i.matchStart != -1){
				num++;
			}
		});
		return num;
	}
		
	$: filterFiles = (files) => {
		files.forEach(function (i) {
			i.tags = filter(i.tags);
		});
		files.sort(function(a, b) {
			console.log(`a.matchs ${numMatches(a.tags)} b.matchs ${numMatches(a.tags)}`);
			if (numMatches(b.tags) > numMatches(a.tags)){
				console.log('1 (CHANGE)');
				return 1;
			}
			else if (numMatches(a.tags) > numMatches(b.tags))
			{
				return -1;
			}
			if (b.name < a.name){
				console.log('1 (CHANGE)');
				return 1;
			}
			console.log('-1 (KEEP)');
			return -1;
		});
		return files;
	}

	$: filter = (tags) => {
		console.log(`sorting tags: ${JSON.stringify(tags)}`);
		tags.forEach(function (t) {
			t.matchStart = searchtext?t.name.indexOf(searchtext):-1;
			console.log(`updated matchStart to ${t.matchStart} inputText:${searchtext}`);
		});
		tags.sort(function(a, b) {
			console.log(`a:${a.name} b:${b.name}`);
			console.log(`a.matchs ${a.matchStart} b.matchs ${b.matchStart}`);
			if (b.matchStart > a.matchStart){
				console.log('1 (CHANGE)');
				return 1;
			}
			else if (a.matchStart > b.matchStart)
			{
				return -1;
			}
			if (b.name < a.name){
				console.log('1 (CHANGE)');
				return 1;
			}
			console.log('-1 (KEEP)');
			return -1;
		});
		console.log(`sorted: ${JSON.stringify(tags)}`);
		return tags;
	}

	const getPreMatchChars = (tag, last) => {
		console.log(`getPreM -> tag: : ${JSON.stringify(tag)} matchStart: ${tag.matchStart}`);
		if (tag.matchStart == -1){
			return tag.name;
		}
			// console.log(`prem: ${tag.name.slice(0, tag.matchStart)}`);
		return tag.name.slice(0, tag.matchStart);
	}
	const getMatchChars = (tag) => {
		console.log(`getmatch: : ${JSON.stringify(tag)} matchStart: ${tag.matchStart}`);
		if (tag.matchStart === -1){
			return '';
		}
			console.log(`match: ${tag.name.slice(tag.matchStart, tag.matchStart+searchtext.length)}`);
		return tag.name.slice(tag.matchStart, tag.matchStart+searchtext.length);
	}
	const getPostMatchChars = (tag) => {
		if (tag.matchStart == -1){
			return ''
		}
		return tag.name.slice(tag.matchStart+searchtext.length, tag.name.length);
	}
	

</script>
<style src='dashboard.scss'>
</style>
<div id='dashboardMain'>
	<TopbarDash/>
	<div id='dashboardContent'>
		<Searchbar bind:searchtext on:search={search}/>
		<div class='dashboardImages'>
			{#each filterFiles(files) as img}
				{#if showFile(img)}
					<div class='dashboardThumb'>
						<img src='{img.data}' alt='image{img.name}'/>
						<div class='dashboardThumbLabel'>
							<div class='dashboardThumbLabelTitle'>
								{img.name}
							</div>
							<div class='dashboardThumbLabelTags'>
								{#each img.tags as tag, n}
									<div class='tagContain' class:notLast={n!==img.tags.length-1}>
										<span class='unmatched tag'>{getPreMatchChars(tag)}</span>
										<span class='matched tag'>{getMatchChars(tag)}</span>
										<span class='unmatched tag'>{getPostMatchChars(tag)}{n===img.tags.length-1?'':','}</span>
									</div>
								{/each}
							</div>
						</div>
					</div>
				{/if}
			{/each}
		</div>
		<div class='noImagesText'>
			{#if hasNoImages}
				Looks like you don’t have any images yet! Get started by clicking ‘upload’ at the top left
			{:else if !loaded}
				Loading...
			{/if}
		</div>
	</div>
</div>
