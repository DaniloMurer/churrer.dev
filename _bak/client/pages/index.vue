<script setup lang="ts">
const showTelemetryAlert = ref(false)

onMounted(async () => {
	const hasMetricsPolicySet = localStorage.getItem('metricsPolicy');
	showTelemetryAlert.value = !hasMetricsPolicySet;
	console.log(typeof hasMetricsPolicySet);
	if (hasMetricsPolicySet === 'true') {
		console.log('sending data')
		sendMetricsData();
	}
})

/**
 * set metric policy in localStorage, based on that the alert is not shown in future visits of the site.
 * furthermore, if the user accepting to the country tracking telemetry, a telemetry call is being made
 * @param metricsPolicy
 */
const setMetricsPolicy = function (metricsPolicy: boolean): void {
	localStorage.setItem('metricsPolicy', String(metricsPolicy));
	//close alert
	showTelemetryAlert.value = false;
	if (metricsPolicy) {
		sendMetricsData();
	} else {
		console.trace('user wont give me their scrumptious data : (');
	}
}

/**
 * send metrics to nuxt backend
 */
const sendMetricsData = function () {
	$fetch('https://hutils.loxal.net/whois').then((whois: any) => {
		$fetch(`/api/telemetry`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json',
			},
			body: {
				'countryName': whois.country,
				'countryIso': whois.countryIso,
				//we split at '.' so the Z part of the ISO string gets truncated, for quarkus and hibernate reasons
				'timestamp': new Date().toISOString()
			}
		}).catch((error) => {
			console.error('Something went wrong. check backend logs....', error);
		});
	}).catch((error) => {
		console.error('Whois doesnt work correctly', error);
	})
}
</script>
<template>
	<div class="flex flex-col items-center gap-16">
		<profile-card />
		<div class="divider">Technologies</div>
		<technology-section />
		<div class="divider">Experience</div>
		<experience-section />
	</div>
	<div role="alert" class="alert fixed bottom-0 left-0 w-2/6 h-56 m-5" v-if="showTelemetryAlert">
		<span class="iconify carbon--information text-2xl"></span>
		<div class="flex flex-row">
			<div class="flex flex-col gap-2 align-bottom">
				<h3 class="font-bold">Lemme grab some of your finest data (respectfully)</h3>
				<div class="text-xs">I only process your country of origin, out of curiosity</div>
				<br>
				<button class="btn btn-sm btn-primary" v-on:click="setMetricsPolicy(true)">Accept</button>
				<button class="btn btn-sm btn-error" v-on:click="setMetricsPolicy(false)">Deny</button>
			</div>
		</div>
	</div>
</template>
