<script>
  import FullScreenFlexPage from '../components/general/fullscreen_flex_page/fullscreen_flex_page.svelte'
	import NewFilePage from '../components/new_file_page/new_file_page.svelte';
	import { onMount } from 'svelte';
	import { sessionToken } from '../store/session.js';
	onMount(async () => {
		if (!$sessionToken){
			window.open('/login', '_self');
		}
	}); 
</script>
<FullScreenFlexPage>
  <NewFilePage/>
</FullScreenFlexPage>
