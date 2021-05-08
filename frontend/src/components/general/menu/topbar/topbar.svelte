<script>
	import ImgHubLogo from '../../logo/logo.svelte'
</script>

<style src='./topbar.scss'></style>

<div id='topbar-container'>
	<div id='topbar'>
		<a id='logotopleft' href='/' target='_self' class='topbar-button'>
			<ImgHubLogo/>
		</a>

		<div class='topbar-group'>
			<a href='/not_implemented' target='_blank' class='topbar-button'>
				Overview
			</a>
			<a href='/not_implemented' target='_blank' class='topbar-button'>
				Demos
			</a>
			<a href='/not_implemented' target='_blank' class='topbar-button'>
				Requirements
			</a>
		</div>
		<div class='topbar-group'>
			<a href='/login' target='_self' class='topbar-button' id='signin'>
				Sign in
			</a>
			<a href='/register' target='_self' class='topbar-button' id='signup'>
				Sign up
			</a>
		</div>
	</div>
</div>
