class ConvaiCommunicator {

	constructor() {
		this.requests = {};
		window.addEventListener("message", (e) => this.receiveMessage(e), false);
	}

	receiveMessage(e) {
		console.log(e);
	}

	async getConfigData() {
		const id = this.makeRand(32);

		const promise = new Promise(((resolve, reject) => {
			this.requests[id] = {resolve, reject};
		}));

		window.opener.postMessage({
			id,
			messageType: "get_config_data",
		});

		return promise;
	}

	async setConfigData(conf) {
		const id = this.makeRand(32);

		const promise = new Promise(((resolve, reject) => {
			this.requests[id] = {resolve, reject};
		}));

		window.opener.postMessage({
			id,
			messageType: "set_config_data",
			data: conf,
		});

		return promise;
	}

	makeRand(length) {
		var result = "";
		var characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
		var charactersLength = characters.length;
		for (var i = 0; i < length; i++) {
			result += characters.charAt(Math.floor(Math.random() * charactersLength));
		}
		return result;
	}
}

const Convai = new ConvaiCommunicator();