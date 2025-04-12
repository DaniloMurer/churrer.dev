export interface Experience {
	id: number;
	position: string;
	company: string;
	timeFrame: string;
	responsibilities: string;
}

export interface Technology {
	id: number,
	name: string;
	logoClass: string;
	experience: string;
	description: string;
}

export class ExperienceDto implements Experience {
	id: number;
	position: string;
	company: string;
	timeFrame: string;
	responsibilities: string;

	constructor(position: string, company: string, timeFrame: string, responsibilities: string) {
		this.id = 0;
		this.position = position;
		this.company = company;
		this.timeFrame = timeFrame;
		this.responsibilities = responsibilities;
	}
}

export class TechnologyDto implements Technology {
	id: number;
	name: string;
	logoClass: string;
	experience: string;
	description: string;

	constructor(name: string, logoClass: string, experience: string, description: string) {
		this.id = 0;
		this.name = name;
		this.logoClass = logoClass;
		this.experience = experience;
		this.description = description;
	}
}
