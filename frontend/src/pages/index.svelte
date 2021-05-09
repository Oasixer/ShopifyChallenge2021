<script>
  import FullScreenFlexPage from '../components/general/fullscreen_flex_page/fullscreen_flex_page.svelte'
  import Homepage from '../components/homepage/homepage.svelte'
	import { sessionToken } from '../store/session.js';
	import { onMount } from 'svelte';
	onMount(async () => {
		if ($sessionToken){
			window.open('/dashboard', '_self');
		}
	}); 
</script>

<FullScreenFlexPage>
	<Homepage/>
</FullScreenFlexPage>
