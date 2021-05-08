<script>
  import FullScreenFlexPage from '../components/general/fullscreen_flex_page/fullscreen_flex_page.svelte'
  import Dashboard from '../components/dashboard/dashboard.svelte'
	import sessionToken from '../store/session';
	import { onMount } from 'svelte';
	import { sessionToken } from '../store/session.js';
	onMount(async () => {
		if (!$sessionToken){
			window.open('/login');
		}
	}); 
</script>

<FullScreenFlexPage>
	{#if $sessionToken}
		<Dashboard/>
	{/if}
</FullScreenFlexPage>
