<script>
	import { onMount } from 'svelte';
	import { sessionToken } from '../store/session.js';
	import Countdown from '../components/general/countdown/countdown.svelte';
	onMount(async () => {
		sessionToken.set('');
	}); 

</script>

<h1>Succesfully logged out. Returning to homepage</h1>
<Countdown countdown={3} on:completed="{() => window.open('/', '_self')}"/>
