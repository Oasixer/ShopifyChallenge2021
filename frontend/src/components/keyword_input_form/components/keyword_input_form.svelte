<script>
	/* import createForm from '@spaceavocado/svelte-form'; */
	/* import TextInput from './input/text.svelte'; */
	import KeywordInput from './keyword_input.svelte';
	import KeywordInputData from './keyword_input_data.js';

	let keywords = [];
	let dummyInvisibleFileInput; // for hacky json file download

	const attachFile = () => {
		dummyInvisibleFileInput.click();
	}
	const import_append_mappings_json = async () => {
		const mappings_text = await dummyInvisibleFileInput.files[fileInput.files.length - 1].text();
		let parsed_keywords = JSON.parse(mappings_text);
		let keywordInputData = new KeywordInputData(parsed_keywords.name,
																								parsed_keywords.aliases,
																								parsed_keywords.related,
																								parsed_keywords.category);
		keywords.push(keywordInputData);
	}

	const clear_all = () => {
		keywords = [];
	}

	const add_keyword_input = () => {
		keywords.push(KeywordInputData.NewInst());
		keywords = [...keywords];
	}
</script>
<style src="./keyword_input_form.scss"/>

<h1>keyword mappings create/edit</h1>

<div class='top_buttons'>
	<button on:click={import_append_mappings_json}>import and append mappings json</button>
	<!-- <button on:click={clear_all}>clear all mappings</button> -->
	<button>export mappings json</button>
</div>
<div class='keywords'>
	{#each Object.values(keywords).sort((a,b) => a.name < b.name || !b.name) as data}
		<KeywordInput {data} all_keywords={keywords}/>
	{/each}
	<button on:click={add_keyword_input}>add keyword input</button>
</div>
