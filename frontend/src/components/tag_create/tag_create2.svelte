<script>
	import Icon from 'svelte-icon'
	import xIcon from '../../../public/assets/icons/x_v2.svg';
	export let tags = [];

	let input = {value: ''}

	const handleTagRemove = (tag) => {
		var index = tags.indexOf(tag);
		if (index !== -1) {
			tags.splice(index, 1);
		}
		tags = [...tags];
	}

	const focusPocus = () => {
		document.getElementById('tagInput').focus();
	}

	const change = () => {
		if (input.value.endsWith(' ') || input.value.endsWith(',')){
				tags = [...tags, {'name': input.value.slice(0,-1)}]; // force reload, add tag
			input.value = '';
			input.ele.value = '';
			focusPocus();
		}
	}

</script>
<style src='tag_create.scss'>
</style>
<div class='tagCreateMain'>
	<div class='tagBar'>
		{#each tags as tag, n}
			<div class='tagContainer top'>
				<span class='unmatched solo tag' >{tag.name}</span>
				<div class='iconContain' on:click={()=>{handleTagRemove(tag)}}>
					<Icon data="{xIcon}" size="14px" stroke="" fill='white'/>
				</div>
			</div>
		{/each}
		<input id='tagInput' type='text' class='tag inputTag' bind:value={input.value} bind:this={input.ele} on:input={change}>
	</div>
</div>
