<script>
	import Section from '../section.svelte';

	let educationInfo = {
		school:'',
		degree:'',
		program:'',
		gradYear:''
	}

</script>

<Section title='Education'>
	<div class='inputs'>
		<h3>School Name</h3>
		<input type='text' bind:value={educationInfo.school}/>
		<h3>Degree Type</h3>
		<input type='text' bind:value={educationInfo.degree}/>
		<h3>Program name</h3>
		<input type='text' bind:value={educationInfo.program}/>
		<h3>(Expected) Graduation Year</h3>
		<input type='text' bind:value={educationInfo.gradYear}/>
	</div>
</Section>
