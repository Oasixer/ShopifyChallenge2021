<script>
	import Section from '../section.svelte';
	import Points from '../../components/points/points.svelte';
	export let title;
	export let workExperiences = [];
	
	const addWorkExperience = () => {
		workExperiences = [
			...workExperiences, 
			{
				jobTitle:'',
				startDate:'',
				endDate:'',
				company:'',
				location:'',
				points:[]
			}
		];
	}
</script>

<Section title={title}>
	<div>
		{#each workExperiences as workExperience}
			<div class='postion'>
				<h3 class='company'>Company</h3>
				<input type='text' bind:value={workExperience.company}/>
				<h3 class='job-title'>Job Title</h3>
				<input type='text' bind:value={workExperience.jobTitle}/>
				<h3 class='date'>Start Date</h3>
				<input type='date' bind:value={workExperience.startDate}/>
				<h3 class='date'>End Date</h3>
				<input type='date' bind:value={workExperience.endDate}/>
				<h3 class='location'>Location</h3>
				<input type='text' bind:value={workExperience.location}/>
				<h3 class='points'>Job Summary</h3>
				<Points bind:points={workExperience.points}/>
			</div>
		{/each}
		<button on:click={addWorkExperience}>
			Add Position
		</button>
	</div>
</Section>
