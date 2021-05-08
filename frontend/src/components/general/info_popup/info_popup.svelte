<script>
	export let title;
	export let message;
	export let bottomLink;
</script>
<style src='info_popup.scss'>
</style>
<div class='popupMain'>
	<div class='popupTitle'>
		{title}
	</div>
	<div class='message'>
		{message}
	</div>
	{#if bottomLink}
		<a href='{bottomLink.href}' target='_self' class='link-button'>
			{bottomLink.title}
		</a>
	{/if}
</div>
