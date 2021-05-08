<script>
	export let points;
	const addPoint = () => {
		points = [...points, {text: '', keywords: []}];
	}

	const addKeyword = (point) => {
		point.keywords = [...point.keywords, ''];
		points = [...points];
	} 
</script>

<style src='./points.scss'>
</style>

<div class='point-section'>
	{#each points as point}
		<div class='point'>
			<input class='point-text' type='text' bind:value={point.text}/>
			<div class='point-keywords'>
				{#each point.keywords as keyword}
					<input type='text' bind:value={keyword}/>
				{/each}
				<button on:click={()=>{addKeyword(point)}}>
					Add Keyword
				</button>  
			</div>
 
		</div>
	{/each}
	<button on:click={addPoint}>
		Add Point
	</button>
</div>
