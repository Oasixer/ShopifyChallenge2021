<script>
	import Section from '../section.svelte';
	import Points from '../../components/points/points.svelte'; 

	export let awards = [];
</script>

<Section title='Award'>
	<Points bind:points={awards}/>
</Section>
