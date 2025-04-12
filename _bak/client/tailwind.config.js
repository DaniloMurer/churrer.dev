const {addIconSelectors} = require('@iconify/tailwind')
module.exports = {
	plugins: [
		require('daisyui'),
		addIconSelectors(['logos', 'carbon'])
	],
	// since tailwind purges unused classes, the dynamic use of inconify classes (meaning I gather them from the server)
	// is an issue. Therefore, the classes I want to use need to be in the safelist so tailwind doesn't purge them when
	// creating the bundle
	safelist : [
		"logos--java",
		"logos--javascript",
		"logos--typescript-icon",
		"logos--kubernetes"
	],
	daisyui: {
		themes: ['night', 'winter']
	},
	theme: {
		extend: {
			animation: {
				'bounce-slow': 'bounce 2s infinite ease-in-out',
			}
		}
	}
};
