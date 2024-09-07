export default async function repost(agent, uri, cid) {
    await agent.repost(uri, cid);
}
