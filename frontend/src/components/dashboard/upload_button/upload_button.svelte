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
					background: 'linear-gradient(80.44deg, #1A0841 11.94%, #580E7B 86.56%)',
					borderRadius: '9px',
					borderBottomRightRadius: '20px',
					borderBottomLeftRadius: '20px'
				},
				styleContent: {
					overflow: 'hidden',
					padding: '0'
				},
				styleCloseButton: {
						top: '-10px',
						right: '-10px',
						border: 'none',
						cursor: 'pointer',
						boxShadow: 'none'
				},
		});
	}
</script>
<style src='upload_button.scss'>
</style>
<button on:click={showUploadPopup}>
	Upload
</button>
