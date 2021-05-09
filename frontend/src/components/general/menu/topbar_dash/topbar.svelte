<script>
	import ImgHubLogo from '../../logo/logo.svelte'
	import UploadButton from '../../../dashboard/upload_button/upload_button.svelte';
</script>

<style src='./topbar.scss'>
</style>

<div id='topbar-container'>
	<div id='topbar'>
		<div class='topbar-group first'>
			<a id='logotopleft' href='/' target='_self' class='topbar-button'>
				<ImgHubLogo/>
			</a>
			<UploadButton/>
		</div>

		<div class='topbar-group'>
			<a href='/dashboard' target='_self' class='topbar-button'>
				Dashboard
			</a>
			<a href='/not_implemented' target='_self' class='topbar-button'>
				Browse
			</a>
		</div>
		<div class='topbar-group'>
			<a href='/logout' target='_self' class='topbar-button' id='signout'>
				Sign out
			</a>
		</div>
	</div>
</div>
