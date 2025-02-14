<script setup lang="ts">
defineProps(['data', 'tableTitle'])
</script>
<template>
	<h1 class="font-bold">{{ tableTitle }}</h1>
	<div>
		<table class="table table-zebra">
			<thead>
				<tr>
					<slot name="table-header" />
				</tr>
			</thead>
			<tbody>
				<tr v-for="item in data">
					<slot name="table-row" :item="item" />
				</tr>
			</tbody>
		</table>
	</div>
</template>
