<script>
	import LoginFields from '../general/form/login_fields/login_fields.svelte';
</script>
<style src='login_page.scss'></style>
<div id=loginPageMain>
	<h1 id='loginTitle'>Login to ImgHub</h1>
	<div id='newHere'>
		<div id='newHereInner'>
			New here?
		</div>
		<a href='/login' target='_self'>Create an account</a>
	</div>
	<div id='loginMain'>
		<LoginFields/>
	</div>
</div>
