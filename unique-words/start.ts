import express from 'express';

interface Parameters {
	text: string
}

function start() {
	const app = express();
	
	app.use(express.json());

	app.get("/healthz", (_req, res) => {
		res.send({
			ok: true
		});
	})
	
	// Register count words controller
	app.post("/op", (req, res) => {
		// Validate parametes
		const body = req.body;

		if (!body || typeof body["text"] !== "string") {
			res.sendStatus(400);
			return;
		}

		const params = body as Parameters;

		// Split string by spaces to find words
		const wordArray = params.text.split(" ");
		const words = wordArray.length;

		res.send({
			key: "unique_words",
			value: words,
			cache_hit: false
		});
	});

	app.listen(8085);
}

start();
