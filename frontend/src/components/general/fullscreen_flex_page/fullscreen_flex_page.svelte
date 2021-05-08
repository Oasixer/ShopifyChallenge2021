<script>
	import { modal } from '../../../store/modal.js';
	import Modal from '../modal/modal.svelte';
</script>
<style src='fullscreen_flex_page.scss'></style>

	<main class='fullscreen-flex-page'>
		<Modal show={$modal}>
			<slot>
			</slot>
		</Modal>
	</main>
