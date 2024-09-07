import express from "express";

export default async function server(): Promise<void> {
  const app = express();

  app.get("/health", (req, res) => {
    res.status(200).send("OK");
  });

  const port = process.env.PORT || 3000;
  app.listen(port, () => {
    console.log(`Health check running on port ${port}`);
  });
}
