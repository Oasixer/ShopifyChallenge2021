<script>
	import ImgHubLogo from '../../logo/logo.svelte';
	import { createEventDispatcher, onDestroy } from 'svelte';
	import { fade, fly } from 'svelte/transition';
	export let modalWidth;
	import Icon from 'svelte-icon';
	import homeIcon from '../../../../../public/assets/icons/home.svg';
	import resumeIcon from '../../../../../public/assets/icons/resume.svg';
	import coverLetterIcon from '../../../../../public/assets/icons/cover_letter.svg';
	import applicationsIcon from '../../../../../public/assets/icons/applications.svg';
	import OpenCloseIcon from '../open_close_icon/open_close_icon.svelte';
	
	export let visible = true;

	// const dispatch = createEventDispatcher();
	// export const close = () => dispatch('close');
	

	const toggle = () => {
		// TODO some animation
		visible = !visible;
	}

	let modal;

	let sidebarButtons = [
		{'name': 'Home', 'href': '/', 'icon': homeIcon, 'size': '30px', 'pad': '18px 3px 0 5px'},
		{'name': 'Resumes', 'href': '/not_implemented', 'icon': resumeIcon, 'size': '30px', 'pad': '20px 3px 0 5px'},
		{'name': 'Cover Letters', 'href': '/not_implemented', 'icon': coverLetterIcon, 'size': '30px', 'pad': '20px 3px 0 4px'},
		{'name': 'Applications', 'href': '/applications', 'icon': applicationsIcon, 'size': '21px', 'pad': '6px 7px 0 9px'}
	];

	const handle_keydown = e => {
		if (e.key === 'Escape') {
			close();
			return;
		}

		if (e.key === 'Tab') {
			// trap focus
			const nodes = modal.querySelectorAll('*');
			const tabbable = Array.from(nodes).filter(n => n.tabIndex >= 0);

			let index = tabbable.indexOf(document.activeElement);
			if (index === -1 && e.shiftKey) index = 0;

			index += tabbable.length + (e.shiftKey ? -1 : 1);
			index %= tabbable.length;

			tabbable[index].focus();
			e.preventDefault();
		}
	};

	const previously_focused = typeof document !== 'undefined' && document.activeElement;

	if (previously_focused) {
		onDestroy(() => {
			previously_focused.focus();
		});
	}
	
	const flyIn = {x: -200, duration: 200};
	const flyOut = {x: -200, duration: 200};
</script>
<style src='./sidebar.scss'>
</style>

<svelte:window on:keydown={handle_keydown}/>

{#if visible}
	<div class="modal-background"
			 bind:clientWidth={modalWidth}
			 in:fly="{flyIn}"
			 out:fly="{flyOut}">
	</div>

	<div
		class="modal"
		role="dialog"
		aria-modal="true"
		bind:this={modal}
		in:fly="{flyIn}"
		out:fly="{flyOut}">
		<div id='sidebarContentContainer'>
			<a id='logo' href='/' target='_self' class='sidebar-button'>
				<ImgHubLogo/>
			</a>
			{#each sidebarButtons as sidebarButton}
				<a href={sidebarButton.href} target='_self' class='sidebar-button'>
					<div class="icon-wrapper" style="padding: {sidebarButton.pad}">
						<Icon data="{sidebarButton.icon}" size="{sidebarButton.size}" height=40px stroke="" color="white" fill="white"/>
					</div>
					<div class='sidebar-text'>
						{sidebarButton.name}
					</div>
				</a>
			{/each}
			<OpenCloseIcon open={true} on:click={toggle}/>
		</div>
	</div>
{:else}
	<div id=fullHeightContainer>
		<OpenCloseIcon open={false} on:click={toggle}/>
	</div>
{/if}
