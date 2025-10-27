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

		// Create map to store unique characters
		const text = params.text;
		const charSet = new Set<string>();

		for (let i = 0; i < text.length; i++) {
			const char = text.charAt(i);
			charSet.add(char);
		}

		res.send({
			key: "unique_chars",
			value: charSet.size,
			cache_hit: false
		});
	});

	app.listen(8085);
}

start();
