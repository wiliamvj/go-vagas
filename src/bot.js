import { AtpAgent } from "@atproto/api";
import * as process from "process";
import WebSocket from "ws";
import { CarReader } from "@ipld/car/reader";
import { cborDecodeMulti, cborDecode } from "@atproto/common";
export default async function bot() {
    const ws = new WebSocket("wss://bsky.network/xrpc/com.atproto.sync.subscribeRepos");
    const agent = new AtpAgent({
        service: "https://bsky.social",
    });
    await agent.login({
        identifier: process.env.BLUESKY_USERNAME,
        password: process.env.BLUESKY_PASSWORD,
    });
    ws.on("message", async (data) => {
        const [header, payload] = cborDecodeMulti(data);
        if (header.op === 1) {
            const t = header?.t;
            if (t) {
                const { ops, blocks } = payload;
                if (ops) {
                    const [op] = ops;
                    if (op?.action === "create") {
                        const cr = await CarReader.fromBytes(blocks);
                        const block = await cr.get(op.cid);
                        const post = cborDecode(block?.bytes || new Uint8Array());
                        if (!post?.text && !post?.reply)
                            return;
                        const terms = [
                            "#govagas",
                            "#golangvagas",
                            "#vagasgolang",
                            "#vagasgo",
                            "#gojobs",
                        ];
                        if (!terms.some((term) => post.text.toLowerCase().includes(term))) {
                            return;
                        }
                        const postUri = post?.reply?.root?.uri;
                        const postCid = post?.reply?.root?.cid;
                        try {
                            // like original post
                            await agent.like(postUri, postCid);
                            // repost original post
                            await agent.repost(postUri, postCid);
                        }
                        catch (e) {
                            console.error("Error", e);
                        }
                    }
                }
            }
        }
    });
}
