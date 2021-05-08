<script>
  import ClickOutside from 'svelte-click-outside';
	export let items;
	// ACCEPTED ITEMS FORMATS:
	// -----FLAT:
	// [{'id':'item1', 'title': 'Item 1'}, ...]
	// -----GROUPED:
	// [{'id':'group1', 'title': 'Group 1', 'items': [{'id':'item1', 'title': 'Item 1'}]}, ...]
	export let label; // ie. 'select resume'
	export let selected; // selected item id (binding)
	export let expanded = false;

	let active = expanded?true:false;

	const dropdown = !expanded;
	const grouped = items[0].hasOwnProperty('items');
	
	let selectedTitle = label;

	let selectWrapperEle; // element binding for click outside

	const select = (item) => {
		selected = item.id;
		selectedTitle = item.title;
		deselectAll();
		item.selected = true;
		active = expanded?true:false;
		items = [...items]; // classic svelte hack to force reload
	}

	const deselectAll = () => {
		items.forEach(function (i) {
			if (grouped){
				i.items.forEach(function (j) {
					j.selected = false;
				});
			}
			else{
				i.selected = false;
			}
		});
	}

	const isSelected = (item) => {
		return item.id === selected;
	}
</script>
<style src='select_box.scss'>
</style>
<div class='select' class:expanded class:dropdown bind:this={selectWrapperEle}>
	<select bind:value={selected}
					class="select-hidden">
		<option value="" disabled selected hidden>
			{selectedTitle}
		</option>
		{#if grouped}
			{#each items as group}
				<optgroup label="{group.title}">
					{#each group.items as item}
						<option class='item indented' value={item.id}>
							{item.title}
						</option>
					{/each}
				</optgroup>
			{/each}
		{:else}
			{#each items as item}
				<option class='item' value={item.id}>
					{item.title}
				</option>
			{/each}
		{/if}
	</select>
	{#if !expanded}
		<div class='select-styled'
				 on:click={()=>{active=!active;}}
				 class:active>
			{selectedTitle}
		</div>
	{/if}
	{#if active}
		<!-- Close the dropdown when you click outside it (I would use on:focusout but it didnt work for some reason) -->
		<!-- We have to exclude the select element itself because it appears outside of the dropdown content -->
		<ClickOutside on:clickoutside={()=>{active=expanded?true:false}} exclude={[selectWrapperEle]}>
			<ul class='select-options' class:expanded class:dropdown>
				{#if !grouped}
					{#each items as item}
						<li class='option' on:click={()=>{select(item)}}
								class:selected={item.selected}>
							{item.title}
						</li>
					{/each}
				{:else}
					{#each items as group}
						<li class='groupTitle'>
							{group.title}
						</li>
						{#each group.items as subItem}
							<li class='option' on:click={()=>{select(subItem)}}
								class:selected={subItem.selected}>
								{subItem.title}
							</li>
						{/each}
					{/each}
				{/if}
			</ul>
		</ClickOutside>
	{/if}
</div>
