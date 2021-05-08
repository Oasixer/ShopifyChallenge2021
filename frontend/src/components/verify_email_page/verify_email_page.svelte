<script>
	import Topbar from '../general/menu/topbar/topbar.svelte';
	const resendVerificationEmail = () => {};
</script>

<style src='./verify_email_page.scss'>
</style>

<div id='verify-page-main'>
	<Topbar/>
	<div id='verify-page-content'>
		<div id='thanks'>
			Thanks for joining ImgHub!
		</div>
		<div id='message'>
			We've sent you a verification email at {"<email>"}, please check your inbox and click on the link
		</div>
		<button id='resend' on:click={resendVerificationEmail}>
			Resend verification email
		</button>
	</div>
</div>
