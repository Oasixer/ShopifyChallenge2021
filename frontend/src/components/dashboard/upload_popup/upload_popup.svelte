<script>
	import Icon from 'svelte-icon';
	import imageIcon from '../../../../public/assets/icons/image.svg';
	let fileInput;
</script>
<style src='upload_popup.scss'>
</style>
<div class=uploadPopup>
	<input type='file' bind:this={fileInput} style='visibility: hidden'>
	<div class='dropImage'>
		Drop image here
	</div>
	<div class='icon-wrapper' on:click={()=>fileInput.click()}>
		<Icon data="{imageIcon}" size="50px" viewBox="0 -10 60 80" stroke=""/>
		<div style='margin: 0 10px'></div>
		Choose photo
	</div>
</div>
