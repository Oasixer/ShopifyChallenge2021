<script>
  import { createEventDispatcher } from 'svelte';
  const dispatch = createEventDispatcher();
  const click = () => dispatch('click');
	const height = 40;
	let width = height/2;
	export let open;
	$: sideMotion = open ? -width: width;
	$: closed = !open;
</script>
<style src='open_close_icon.scss'>
</style>
	
<div id='iconContainer'
	on:click>
	<svg
		width="{width}px"
		height="{height}px"
		class:closed>
		<path 
				id="svg_4"
				d="m{open?width:0},0 l0,{height} l{sideMotion},{-height/2} l{-sideMotion},{-height/2}"
				opacity="NaN"
				fill-opacity="1"
				fill="white"
				stroke-linecap="round"
				/>
	</svg>
</div>
