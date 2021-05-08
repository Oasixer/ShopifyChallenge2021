<script>
	import Points from '../../components/points/points.svelte'
	import Section from '../section.svelte'
	export let interests = []
</script>
<style src='interests.scss'>
</style>

<Section title='Interests'>
	<Points bind:points={interests}/>
</Section>
