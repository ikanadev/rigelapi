import http from "k6/http";

export const options = {
	// vus: 5,
	// duration: '10s',
	scenarios: {
		low: {
			executor: "constant-arrival-rate",
			maxVUs: 1000,
			preAllocatedVUs: 25,
			rate: 480,
			timeUnit: "1m",
			duration: "1m",
			gracefulStop: "2m",
		},
	},
};

export default function () {
	// http.get("https://api.auleca.com/deps");
	http.get("https://api.auleca.com/provs/dep/2");

  /*
	const payload = JSON.stringify({
		email: "kevin@gmail.com",
		password: "troyalk7727",
	});
	const params = {
		headers: {
			"Content-Type": "application/json",
		},
	};
	http.post("https://api.auleca.com/signin", payload, params);
  */

	// http.get("https://api.auleca.com/class/6k5jNCOtjkiACHIB49jQP");
}
