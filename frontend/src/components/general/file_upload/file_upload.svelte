<script>
	import ToggleList from '../form/toggle_list/toggle_list.svelte';
	import SelectBox from '../form/select_box/select_box.svelte';

	let show_generated_resumes = true;
	let show_uploaded_resumes = true;
	let toggle_items = [
		{'title': 'Show generated resumes', 'on': show_generated_resumes},
		{'title': 'Show uploaded resumes', 'on': show_uploaded_resumes}
	];
		
	let select_items_flat_short = [
		{'title': 'option 1', 'id': 'option1'},
		{'title': 'option 2', 'id': 'option2'}
	];
	let select_items_flat_long = [
		{'title': 'option 1', 'id': 'option1'},
		{'title': 'option 2', 'id': 'option2'},
		{'title': 'option 3', 'id': 'option3'},
		{'title': 'option 4', 'id': 'option4'},
		{'title': 'option 5', 'id': 'option5'},
		{'title': 'option 6', 'id': 'option6'},
		{'title': 'option 7', 'id': 'option7'},
		{'title': 'option 8', 'id': 'option8'}
	];
	let select_items_grouped_short = [
		{'id': 'group1', 'title': 'Group 1', 'items': [
			{'title': 'option 1', 'id': 'option1'},
			{'title': 'option 2', 'id': 'option2'}
			]},
		{'id': 'group2', 'title': 'Group 2', 'items': [
			{'title': 'option 1', 'id': 'option1'},
			{'title': 'option 2', 'id': 'option2'}
		]}
	];
		
	let select_items_grouped_long = [
		{'id': 'group1', 'title': 'Group 1', 'items': [
			{'title': 'option 1', 'id': 'option1'},
			{'title': 'option 2', 'id': 'option2'},
			{'title': 'option 3', 'id': 'option3'},
			{'title': 'option 4', 'id': 'option4'},
			{'title': 'option 5', 'id': 'option5'}
			]},
		{'id': 'group2', 'title': 'Group 2', 'items': [
			{'title': 'option 1', 'id': 'option21'},
			{'title': 'option 2', 'id': 'option22'}
		]}
	];
</script>

<style src='file_upload.scss'>
</style>

<div class='file_upload'>
	<h1>Your resume PDFs</h1>
	<ToggleList bind:items={toggle_items}/>
	<div style='display: flex; flex-flow: column nowrap; align-items: stretch; width: 100%;'>
		<SelectBox
			bind:items={select_items_flat_short}
			expanded={false}
			label={'select from flat_short'}/>
		<SelectBox
			bind:items={select_items_grouped_short}
			expanded={false}
			label={'select from grouped_short'}/>
		<SelectBox
			bind:items={select_items_flat_long}
			expanded={false}
			label={'select from flat_long'}/>
		<SelectBox
			bind:items={select_items_grouped_long}
			expanded={false}
			label={'select from grouped_long'}/>
		<SelectBox
			bind:items={select_items_flat_short}
			expanded={true}
			label={'select from flat_short'}/>
		<SelectBox
			bind:items={select_items_grouped_short}
			expanded={true}
			label={'select from grouped_short'}/>
		<SelectBox
			bind:items={select_items_flat_long}
			expanded={true}
			label={'select from flat_long'}/>
		<SelectBox
			bind:items={select_items_grouped_long}
			expanded={true}
			label={'select from grouped_long'}/>
	</div>
</div>
