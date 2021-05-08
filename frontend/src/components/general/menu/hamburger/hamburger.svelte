<script>
  import { createEventDispatcher } from 'svelte';
  const dispatch = createEventDispatcher();
  const click = () => dispatch('click');
</script>
<button on:click aria-label="Menu">
	<svg id="icon" width="184" height="184" viewBox="0 0 184 184" fill="none" xmlns="http://www.w3.org/2000/svg">
		<g id="close">
			<path id="topLine" d="M36 54H148"/>
			<path id="bottomLine" d="M36 130H148"/>
		</g>
		<path id="middleBarCircle" d="M36 92H147.5C161 92 179.5 83.4 179.5 59C179.5 28.5 144 4 92 4C43.3989 4 4 43.3989 4 92C4 140.601 43.3989 180 92 180C140.601 180 180 140.601 180 92C180 43.3989 141.601 4 93 4"/>
	</svg>
</button>
