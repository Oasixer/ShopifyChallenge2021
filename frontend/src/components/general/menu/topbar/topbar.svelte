<script>
	import ImgHubLogo from '../../logo/logo.svelte'
</script>

<style src='./topbar.scss'></style>

<div id='topbar-container'>
	<div id='topbar'>
		<a id='logotopleft' href='/' target='_self' class='topbar-button'>
			<ImgHubLogo/>
		</a>
		<a href='/not_implemented' target='_blank' class='topbar-button' style='margin-right: auto; margin-left: auto;'>
			Overview
		</a>
		<div class='topbar-group'>
			<a href='/login' target='_self' class='topbar-button' id='signin'>
				Sign in
			</a>
			<a href='/register' target='_self' class='topbar-button' id='signup' style='margin-right: 25px; margin-left: 25px;'>
				Sign up
			</a>
		</div>
	</div>
</div>
