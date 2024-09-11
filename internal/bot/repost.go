package bot

import "fmt"

func repost(p *Post) {
	fmt.Println("respost", p)

	if p.Reply != nil {
		fmt.Println("reply", p.Reply.Root.Cid)
	}

}
