<script>
	import Icon from 'svelte-icon'
	import searchIcon from '../../../../../public/assets/icons/search.svg';
	export let searchtext;

  import { createEventDispatcher } from 'svelte';
  
  const dispatch = createEventDispatcher();
		const click = () => {
			dispatch('search');
		}
</script>
<style src='searchbar.scss'>
</style>
<div class='searchOuter'>
	<button class='searchIconWrapper' on:click={click}>
		<Icon data="{searchIcon}" fill="white" size="40" stroke="" viewBox="0 0 24 22"/>
	</button>
	<input type='text' placeholder='Search your images' bind:value={searchtext}/>
</div>
