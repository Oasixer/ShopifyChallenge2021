<script>
	let skills = [];
	import Skills from './sections/skills/skills.svelte';
	import Experience from './sections/experience/experience.svelte';
	import Education from './sections/education/education.svelte';
	import Awards from './sections/awards/awards.svelte';
	import ContactInfo from './sections/contact_info/contact_info.svelte';
	import Interests from './sections/interests/interests.svelte';
	import Sidebar from '../general/menu/sidebar/sidebar.svelte';
	let skillsKeywords;
	let sidebarWidth;
</script>

<style src='./resume_input_form.scss'>
</style>

<Sidebar bind:modalWidth={sidebarWidth}/>
<div id="resumeInputMain" style="padding-left: 300px">
	<h1>Resume Input Form</h1>

	<ContactInfo/>
	<Skills bind:skillsKeywords/>
	<Experience title='Work Experience'/>
	<Experience title='Projects'/>
	<Education/>
	<Awards/>
	<Interests/>
</div>
