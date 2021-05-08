<script>
		import Section from '../section.svelte'
		import LabeledTextInput from '../../../general/form/labeled_text_input/labeled_text_input.svelte'
		let contactInfo = [
				{id: 'firstName', title: 'First Name', value: ''},
				{id: 'lastName', title: 'Last Name', value: ''},
				{id: 'phone', title: 'Phone Number', value: ''},
				{id: 'email', title: 'Email Address', value: ''},
				{id: 'linkedIn', title: 'LinkedIn', value: ''},
				{id: 'github', title: 'Github', value: ''},
				{id: 'website', title: 'Website', value: ''}
		]
</script>

<style src='contact_info.scss'>
</style>

<Section title='Contact Information'>
		<div class='inputs'>
				{#each contactInfo as item}
						<LabeledTextInput bind:value={item.value} label="{item.title}"/>
				{/each}
		</div>
</Section>
