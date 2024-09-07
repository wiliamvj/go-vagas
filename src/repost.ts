import { AtpAgent } from "@atproto/api";

export default async function repost(
  agent: AtpAgent,
  uri: string,
  cid: string
): Promise<void> {
  await agent.repost(uri, cid);
}
