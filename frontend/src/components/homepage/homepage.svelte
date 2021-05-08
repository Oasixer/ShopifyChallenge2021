<script>
	import TopBar from '../general/menu/topbar/topbar.svelte';
	import {email} from '../../store/register';

	let emailVal='';
	const setEmailStore = () => {
		email.set(emailVal);
	}
</script>
<style src='./homepage.scss'>
</style>
<div id='landing-page-main'>
	<TopBar/>
	<div id="landing-page-content">
		<div id="landing-page-horizontal-split">
			<div id="landing-page-text">
				<div id="start-uploading-today">
					<div id="start-uploading">
						Start uploading
					</div>
					<div id="today">
						today
					</div>
				</div>
				<div id='lower-text'>
					ImgHub is a secure image hosting site that puts privacy first and also please hire me shopify
				</div>
				<div id='register-bar'>
					<input type="text" bind:value={emailVal}
					on:change={setEmailStore}/>
					<a href="/register" id="register-button" target="_self">
						Register
					</a>
				</div>
			</div>
			<div id='demo-animation-placeholder'>
				demo animation coming later
			</div>
		</div>
	</div>
</div>
