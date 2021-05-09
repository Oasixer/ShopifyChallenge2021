<script>
	import UploadPopup from '../upload_popup/upload_popup.svelte';
  import { getContext } from 'svelte';
  const { open } = getContext('popup-modal');
	const showUploadPopup = () => {
		open(UploadPopup,
			{},
			{
			styleWindowWrap: {
				margin: '0'
			},
			styleWindow: {
				border: '0',
				background: 'linear-gradient(80.44deg, #1A0841 11.94%, #580E7B 86.56%)'
			},
			styleContent: {
			},
			styleCloseButton: {
					position: 'relative',
					top: '-20px'
			},
		});
	}
</script>
<style src='upload_button.scss'>
</style>
<button on:click={showUploadPopup}>
	Upload
</button>
