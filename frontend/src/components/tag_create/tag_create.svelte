<script>
	import Icon from 'svelte-icon'
	import xIcon from '../../../public/assets/icons/x_v2.svg';
	let allTags = ['sampleTag', 'jeff', 'birds', 'home'].map(x => {return {'name': x, 'matchStart': -1}});
	export let tags = [{'name': '', 'ele': null}];
	const handleTagRemove = (tag) => {
		var index = tags.indexOf(tag);
		if (index !== -1) {
			tags.splice(index, 1);
		}
		tags = tags.filter(x=>{x.name.length});
		allTags = [...allTags, tag];
		let newTag = {'name': '', 'ele': null};
		focused = newTag;
		tags = [...tags, newTag]; // force reload
		focusPocus();
	};
	const getMatchChars = (tag) => {
		console.log(`tag: ${JSON.stringify(tag)} matchStart: ${tag.matchStart}`);
		if (tag.matchStart === -1){
			return ''
		}
			console.log(`match: ${tag.name.slice(tag.matchStart, tag.matchStart+focused.name.length)}`);
		return tag.name.slice(tag.matchStart, tag.matchStart+focused.name.length);
	}
	const getPostMatchChars = (tag) => {
		if (tag.matchStart == -1){
			return ''
		}
		return tag.name.slice(tag.matchStart+focused.name.length, tag.name.length);
	}
	const focusPocus = async () => {
		setTimeout(()=>{
					console.log(`focused.name: ${focused.name}`);
					if (!focused.ele){
						focused.ele = document.getElementById(`_${focused.name}`);
					}
					focused.ele.focus();
					if (tags.indexOf(focused) === tags.length-1){
						focused.ele.style.width = 'auto';
					}
					else{
						focused.ele.size = Math.max(focused.name.length, 5);
					}
				},400);
	}
	import { onMount } from 'svelte';
	onMount(async () => {
		await focusPocus();
		focused.ele.style.width = '100%';
	}); 

	const focusout = (tag, n) => {
		// if (tag.name)
		// if (tags.length===1){
		// }
	}
	let focused = tags[0];
	$: console.log(`focused: ${JSON.stringify(focused)}`);
	$: console.log(`tags: ${JSON.stringify(tags)}`);

	const filter = (alltags, tags) => {
		console.log(`sorting alltags: ${JSON.stringify(alltags)}`);
		alltags.forEach(function (t) {
			t.matchStart = focused.name?t.name.indexOf(focused.name):-1;
				console.log(`updated matchStart to ${t.matchStart} inputText:${focused.name}`);
		});
		let tagnames = tags.map(x => x.name);
		alltags = alltags.filter(x=>!(tagnames.includes(x.name)&&focused.name!==x.name));
		alltags.sort(function(a, b) {
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
		console.log(`sorted: ${JSON.stringify(alltags)}`);
		return alltags;
	}
	const getPreMatchChars = (tag) => {
			console.log(`getPreM: ${JSON.stringify(tag)}`);
		if (tag.matchStart == -1){
			return tag.name
		}
			// console.log(`prem: ${tag.name.slice(0, tag.matchStart)}`);
		return tag.name.slice(0, tag.matchStart);
	}
	const change = (tag, n) => {
		if (n === tags.length-1){
			tag.ele.style.width = '100%';
		}
		else{
			tag.ele.size = Math.max(tag.name.length, 5);
		}
		if (tag.name.endsWith(',') || tag.name.endsWith(',')){
			tag.name = tag.name.slice(0,-1);
			let newTag = {'name': '', 'ele': null};
			tags = [...tags, newTag]; // force reload
			focused = newTag;
			focusPocus();
		}
		
		allTags = [...filter(allTags, tags)];
	}
	const handleTagClick = (tag) => {
		focused.ele.value=`${tag.name},`;
		let event = new Event('input', {
			bubbles: true,
			cancelable: true,
		});
		focused.ele.dispatchEvent(event);	 
	}
	// $: filter(allTags);
</script>
<style src='tag_create.scss'>
</style>
<div class='tagCreateMain'>
	<div class='tagTopBar tagBar'>
		{#each tags as tag, n}
			{#if focused==tag}
				<input id='_{tag.name}' type='text' class='tag inputTag' bind:value={tag.name} on:focusout={()=>{focusout(tag, n)}} bind:this={tag.ele} on:input={()=>{change(tag,n)}}>

				{:else}
					<div class='tagContainer top'>
						<span class='unmatched solo tag' on:click={()=>{console.log('12345');focused=tag; tags = [...tags]; focusPocus();}}>{tag.name}</span>
						<div class='iconContain' on:click={()=>{handleTagRemove(tag)}}>
							<Icon data="{xIcon}" size="14px" stroke="" fill='white'/>
						</div>
					</div>
				{/if}
		{/each}
	</div>
	<div class='line'></div>
	<div class='tagBotBar tagBar'>
		{#each allTags as tag}
			<div class='tagContainer bot' on:click={()=>{handleTagClick(tag)}}>
				<span class='unmatched tag'>{getPreMatchChars(tag)}</span>
				<span class='matched tag'>{getMatchChars(tag)}</span>
				<span class='unmatched tag'>{getPostMatchChars(tag)}</span>
			</div>
		{/each}
	</div>
</div>
