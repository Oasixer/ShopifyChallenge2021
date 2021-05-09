<script>
	import InfoPopup from '../info_popup/info_popup.svelte'
  import { getContext } from 'svelte';
  const { open } = getContext('popup-modal');
	const showPopup = () => {
		open(InfoPopup, {message: 'All your images are private, visible only to you. Click on an image to view full-size, download, edit tags, or delete it', title: 'Help'});
	}
	// import {createEventDispatcher} from 'svelte';
	// const dispatch = createEventDispatcher();
	// export const click = () => dispatch('click');
</script>
<style src='help_icon.scss'>
</style>
<div class='helpIcon' on:click={showPopup}>
	?
</div>
