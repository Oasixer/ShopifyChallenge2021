<script>
	import Section from '../section.svelte';
	let open = null;
	export let skillsSections = [];
	const expandKeywordList = () => {
		open = 'keywordList';
	}

	const expandSentenceList = () => {
		open = 'sentenceList';
	}

	const addSkillSection = () => {
		skillsSections = [...skillsSections, {name: '', keywords: []}];
	}
</script>

<style src='./skills.scss'></style>

<Section title='Skills'>
	{#if !open}
		<div class='skill-buttons'>
			<button on:click={expandKeywordList}>
				Keyword List
			</button>
		
			<button on:click={expandSentenceList}>
				Sentence List
			</button>
		</div>
	{/if}

	
	<div class='inputs'>
		{#if open == 'keywordList'}
			<div class='keyword-inputs'>
				{#each skillsSections as skillsSection}
					<div class='title'>Section:</div>
					<input class='skills-section' type='text' bind:value={skillsSection.name}/>
					<div class='title'>Keywords:</div>
					{#each skillsSection.keywords as keyword}
						<input class='skills-keyword' type='text' bind:value={keyword}/>
					{/each}
					<button on:click={()=>{skillsSection.keywords = [...skillsSection.keywords, '']}}>
						Add Keyword
					</button>
				{/each}
			</div>
			<button on:click={addSkillSection}>
				Add Skills Section
			</button>
		{/if}
		{#if open == 'sentenceList'}
			<div class='keyword-inputs'>
				{#each skillsSections as skillsSection}
					<div class='title'>Sentence:</div>
					<input class='skills-section' type='text' bind:value={skillsSection.name}/>
					<div class='title'>Keywords:</div>
					{#each skillsSection.keywords as keyword}
						<input class='skills-keyword' type='text' bind:value={keyword}/>
					{/each}
					<button on:click={()=>{skillsSection.keywords = [...skillsSection.keywords, '']}}>
						Add Keyword
					</button>
				{/each}
			</div>
			<button on:click={addSkillSection}>
				Add Skills Section
			</button>
		{/if}
		

	</div>

</Section>
